package impl

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"pet-paradise/log"
	"pet-paradise/middleware"
	"pet-paradise/model"
	"pet-paradise/utils"
	"regexp"
	"strconv"
)

const PASSWORD_REG = `^[a-zA-Z]\w{7,15}$`

func Login(ctx *gin.Context) {
	log.Logger().Info("[Login] %s", ctx.ClientIP())

	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")
	userInfo, err := model.UserTable.GetOneByName(userName)

	if err == sql.ErrNoRows {
		utils.Response(ctx, http.StatusUnauthorized, "there's no such user："+userName, nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)); err != nil {
		utils.Fail(ctx, "wrong password", nil)
		return
	}

	token, err := middleware.ReleaseToken(userInfo.ID, userInfo.Role, ctx.ClientIP())
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "succeed", gin.H{"token": token})
}

func Logout(ctx *gin.Context) {
	log.Logger().Info("[Logout] %s", ctx.ClientIP())
	utils.Success(ctx, "ok", nil)
}

func Register(ctx *gin.Context) {
	log.Logger().Info("[Register] %s", ctx.ClientIP())

	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if ok, _ := regexp.MatchString(PASSWORD_REG, password); !ok {
		utils.Fail(ctx, "password must begin with letters, between 8-16 in length, can only contain letters, numbers and underscores", nil)
		return
	}

	if info, err := model.UserTable.GetOneByName(userName); err == sql.ErrNoRows {
	} else if err != nil {
		utils.Fail(ctx, err.Error(), nil)
		return
	} else if info.Name == userName {
		utils.Fail(ctx, "user has been registered："+userName, nil)
		return
	}

	var userInfo model.UserInfo
	if err := ctx.Bind(&userInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	userInfo.Password = string(hashedPassword)
	if _, err := model.UserTable.InsertNewUserInfo(userInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func GetUserInfo(ctx *gin.Context) {
	log.Logger().Info("[GetUserInfo] %s", ctx.ClientIP())

	userID := ctx.GetString("uid")

	userInfo, err := model.UserTable.GetOneById(userID)
	if err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
	}

	utils.Success(ctx, "ok", userInfo)
}

func UpdateUserInfo(ctx *gin.Context) {
	log.Logger().Info("[UpdateUserInfo] %s", ctx.ClientIP())

	userID := ctx.GetString("uid")
	var userNewInfo = model.UserInfo{}
	if err := ctx.Bind(&userNewInfo); err != nil {
		utils.Fail(ctx, "invalid param", nil)
		return
	}

	if userNewInfo.Password != "" {
		userNewInfo.Password = ""
	}
	if userNewInfo.Role != "" {
		userNewInfo.Role = ""
	}

	if _, err := model.UserTable.UpdateUserInfoById(userNewInfo, userID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
	utils.Success(ctx, "ok", nil)
}

func UpdateUserPassword(ctx *gin.Context) {
	log.Logger().Info("[UpdateUserPassword] %s", ctx.ClientIP())

	userID := ctx.GetString("uid")
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")

	if ok, _ := regexp.MatchString(PASSWORD_REG, newPassword); !ok {
		utils.Fail(ctx, "password must begin with letters, between 8-16 in length, can only contain letters, numbers and underscores", nil)
		return
	}

	if oldPassword == newPassword {
		utils.Fail(ctx, "this password has been used recently", nil)
		return
	}

	userInfo, err := model.UserTable.GetOneById(userID)
	if err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(oldPassword)); err != nil {
		utils.Fail(ctx, "密码错误", nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if _, err := model.UserTable.UpdateUserInfoById(model.UserInfo{
		Password: string(hashedPassword),
	}, userID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)

}

func DeleteUser(ctx *gin.Context) {
	log.Logger().Info("[DeleteUser] %s", ctx.ClientIP())

	userID := ctx.PostFormArray("uid")

	if _, err := model.UserTable.DeleteUserInfoById(userID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
	utils.Success(ctx, "ok", nil)
}

func AddAddressInfo(ctx *gin.Context) {
	log.Logger().Info("[AddAddressInfo] %s", ctx.ClientIP())

	userID := ctx.GetString("uid")

	var newAddressInfo model.UserAddressInfo
	if err := ctx.Bind(&newAddressInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}
	newAddressInfo.UserID, _ = strconv.Atoi(userID)
	if _, err := model.AddressTable.InsertNewAddressInfo(newAddressInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func UpdateAddressInfo(ctx *gin.Context) {
	log.Logger().Info("[UpdateAddressInfo] %s", ctx.ClientIP())

	addressID := ctx.PostForm("aid")

	var newAddressInfo model.UserAddressInfo
	if err := ctx.Bind(&newAddressInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if _, err := model.AddressTable.UpdateAddressInfoById(newAddressInfo, addressID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func GetAllAddress(ctx *gin.Context) {
	log.Logger().Info("[GetAllAddress] %s", ctx.ClientIP())

	userID := ctx.GetString("uid")

	addressSlice, err := model.AddressTable.SelectAddressInfoByUserId(userID)
	if err == sql.ErrNoRows {
		utils.Success(ctx, "none", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", err)
		return
	}

	utils.Success(ctx, "ok", addressSlice)
}

func DeleteAddress(ctx *gin.Context) {
	log.Logger().Info("[DeleteAddress] %s", ctx.ClientIP())

	addressIDs := ctx.PostFormArray("aid")

	if _, err := model.AddressTable.DeleteAddressInfoById(addressIDs); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func AdminGetUserInfoByName(ctx *gin.Context) {
	log.Logger().Info("[AdminGetAllUser] %s", ctx.ClientIP())

	userName := ctx.Query("username")
	userInfo, err := model.UserTable.GetOneByName(userName)
	if err == sql.ErrNoRows {
		utils.Success(ctx, "none", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", userInfo)
}

func AdminUpdateUserRole(ctx *gin.Context) {
	log.Logger().Info("[AdminUpdateUserRole] %s", ctx.ClientIP())

	userID := ctx.PostForm("uid")
	newRole := ctx.PostForm("role")

	if _, err := model.UserTable.UpdateUserRoleById(userID, newRole); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}
