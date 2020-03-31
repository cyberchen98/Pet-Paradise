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
	"strconv"
)

type UserInfo struct {
	UserName string
	Password string
}

func Login(ctx *gin.Context) {
	log.Logger().Info("[Login] ", ctx.Request.URL)

	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")
	userInfo, err := model.UserTable.GetOneByName(userName)

	if err == sql.ErrNoRows {
		utils.Response(ctx, http.StatusUnauthorized, "不存在用户："+userName, nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)); err != nil {
		utils.Fail(ctx, "密码错误", nil)
		return
	}

	token, err := middleware.ReleaseToken(userInfo.ID, ctx.ClientIP())
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "登陆成功", gin.H{"token": token})
}

func Logout(ctx *gin.Context) {
	log.Logger().Info("[Logout] ", ctx.Request.URL)
	utils.Success(ctx, "ok", nil)
}

func Register(ctx *gin.Context) {
	log.Logger().Info("[Register] ", ctx.Request.URL)
	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if info, err := model.UserTable.GetOneByName(userName); err == sql.ErrNoRows {
	} else if err != nil {
		utils.Fail(ctx, err.Error(), nil)
		return
	} else if info.Name == userName {
		utils.Fail(ctx, "用户已存在："+userName, nil)
		return
	}

	var userInfo model.UserInfo
	if err := ctx.ShouldBind(&userInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	userInfo.Password = string(hashedPassword)
	if err := model.UserTable.InsertNewUserInfo(userInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func GetUserInfo(ctx *gin.Context) {
	log.Logger().Info("[GetUserInfo] ", ctx.Request.URL)

	userID := ctx.GetInt("user_id")
	userInfo, err := model.UserTable.GetOneById(userID)
	if err != nil {
		return
	}

	utils.Success(ctx, "ok", userInfo)
}

func UpdateUserInfo(ctx *gin.Context) {
	log.Logger().Info("[UpdateUserInfo] ", ctx.Request.URL)

	userID := ctx.GetInt("user_id")
	var userNewInfo = model.UserInfo{}
	if err := ctx.Bind(&userNewInfo); err != nil {
		utils.Fail(ctx, "invalid param", nil)
		return
	}

	if userNewInfo.Password != "" {
		userNewInfo.Password = ""
	}

	if err := model.UserTable.UpdateUserInfoById(userNewInfo, userID); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
	utils.Success(ctx, "ok", nil)
}

func UpdateUserPassword(ctx *gin.Context) {
	log.Logger().Info("[UpdateUserPassword] ", ctx.Request.URL)

	userID := ctx.GetInt("user_id")
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	if oldPassword == newPassword {
		utils.Fail(ctx, "this password has been used recently", nil)
		return
	}

	userInfo, err := model.UserTable.GetOneById(userID)
	if err != nil {
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

	if err := model.UserTable.UpdateUserInfoById(model.UserInfo{
		Password: string(hashedPassword),
	}, userID); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)

}

func DeleteUser(ctx *gin.Context) {
	log.Logger().Info("[DeleteUser] ", ctx.Request.URL)

	userID := ctx.GetInt("user_id")

	if err := model.UserTable.DeleteUserInfoById(userID); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
	utils.Success(ctx, "ok", nil)
}

func AddAddressInfo(ctx *gin.Context) {
	log.Logger().Info("[AddAddressInfo] ", ctx.Request.URL)

	userID := ctx.GetInt("user_id")
	addressInfo := model.UserAddressInfo{
		UserID:      userID,
		Province:    ctx.PostForm("province"),
		City:        ctx.PostForm("city"),
		Details:     ctx.PostForm("details"),
		PhoneNumber: ctx.PostForm("phone_number"),
		Receiver:    ctx.PostForm("receiver"),
		PostCode:    ctx.PostForm("post_code"),
	}

	if err := model.AddressTable.InsertNewAddressInfo(addressInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
}

func UpdateAddressInfo(ctx *gin.Context) {
	log.Logger().Info("[UpdateAddressInfo] ", ctx.Request.URL)

	addressID, err := strconv.Atoi(ctx.PostForm("address_id"))
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	var newAddressInfo model.UserAddressInfo
	if err := ctx.Bind(&newAddressInfo); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	if err := model.AddressTable.UpdateAddressInfoById(newAddressInfo, addressID); err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
}

func GetAllAddress(ctx *gin.Context) {
	log.Logger().Info("[GetAllAddress] ", ctx.Request.URL)

	userID := ctx.GetInt("user_id")

	addressSlice, err := model.AddressTable.GetAllByUserId(userID)
	if err == sql.ErrNoRows {
		utils.Success(ctx, "ok", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", addressSlice)
}