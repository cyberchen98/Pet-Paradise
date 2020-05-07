package impl

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-paradise/log"
	"pet-paradise/model"
	"pet-paradise/utils"
	"strconv"
)

func GetAllOrderInfoByUserId(ctx *gin.Context) {
	log.Logger().Info("[GetAllOrderInfoByUserId] %s", ctx.ClientIP())

	userID := ctx.GetString("uid")

	orderInfoSlice, err := model.OrderTable.SelectOrderInfoByUserId(userID)
	if err == sql.ErrNoRows {
		utils.Success(ctx, "ok", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", orderInfoSlice)
}

func GetOrderInfoById(ctx *gin.Context) {
	log.Logger().Info("[GetOrderInfoById] %s", ctx.ClientIP())

	orderID := ctx.Query("oid")

	orderInfo, err := model.OrderTable.GetOneById(orderID)
	if err == sql.ErrNoRows {
		utils.Fail(ctx, "no record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", orderInfo)
}

func DeleteOrderById(ctx *gin.Context) {
	log.Logger().Info("[DeleteOrderById] %s", ctx.ClientIP())

	orderIDs := ctx.PostFormArray("oid")

	if _, err := model.OrderTable.DeleteOrderInfoById(orderIDs); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func GenerateOrder(ctx *gin.Context) {
	log.Logger().Info("[GenerateOrder] %s", ctx.ClientIP())

	userID, err := strconv.Atoi(ctx.GetString("uid"))
	if err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}
	addressID, err := strconv.Atoi(ctx.PostForm("aid"))
	if err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}
	productID, err := strconv.Atoi(ctx.PostForm("pid"))
	if err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}

	var newOrderInfo model.OrderInfo
	if err := ctx.Bind(&newOrderInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}

	newOrderInfo.UserID = userID
	newOrderInfo.AddressID = addressID
	newOrderInfo.ProductID = productID

	if _, err := model.OrderTable.InsertNewOrderInfo(newOrderInfo); err != nil {
		fmt.Println("err:", err)
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func UpdateOrderInfoById(ctx *gin.Context) {
	log.Logger().Info("[UpdateOrderInfoById] %s", ctx.ClientIP())

	orderID := ctx.PostForm("oid")

	var orderInfo model.OrderInfo
	if err := ctx.Bind(&orderInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}
	if orderInfo.Status != "" {
		orderInfo.Status = ""
	}

	if _, err := model.OrderTable.UpdateOrderInfoById(orderInfo, orderID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func AdminUpdateOrderInfoById(ctx *gin.Context) {
	log.Logger().Info("[AdminUpdateOrderInfoById] %s", ctx.ClientIP())

	orderID := ctx.PostForm("oid")

	var orderInfo model.OrderInfo
	if err := ctx.Bind(&orderInfo); err != nil {
		utils.Fail(ctx, "invalid params", nil)
		return
	}

	if _, err := model.OrderTable.UpdateOrderInfoById(orderInfo, orderID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}
