package impl

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-paradise/log"
	"pet-paradise/model"
	"pet-paradise/utils"
	"strconv"
)

func GetOrderInfoByUserId(ctx *gin.Context) {
	log.Logger().Info("[GetOrderInfoByUserId] ", ctx.Request.URL)

	userID := ctx.GetInt("user_id")

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

func DeleteOrderById(ctx *gin.Context) {
	log.Logger().Info("[DeleteOrderById] ", ctx.Request.URL)

	orderID, err := strconv.Atoi(ctx.PostForm("order_id"))
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}
	if _, err := model.OrderTable.DeleteOrderInfoById(orderID); err == sql.ErrNoRows {
		utils.Fail(ctx, "no this record", nil)
		return
	} else if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	utils.Success(ctx, "ok", nil)
}

func GenerateOrder(ctx *gin.Context) {
	log.Logger().Info("[GenerateOrder] ", ctx.Request.URL)
}

func GetOrdersByProductId(ctx *gin.Context) {
	log.Logger().Info("[GetOrdersByProductId] ", ctx.Request.URL)

}

func UpdateOrderInfoById(ctx *gin.Context) {
	log.Logger().Info("[UpdateOrderInfoById] ", ctx.Request.URL)

	orderID, err := strconv.Atoi(ctx.PostForm("order_id"))
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, "internal error", nil)
		return
	}

	var orderInfo model.OrderInfo
	if err := ctx.Bind(&orderInfo); err != nil {

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
