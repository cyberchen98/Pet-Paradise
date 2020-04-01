package impl

import (
	"github.com/gin-gonic/gin"
	"pet-paradise/log"
)

func GetOrderInfoByUserId(ctx *gin.Context) {
	log.Logger().Info("[GetOrderInfoByUserId] ", ctx.Request.URL)

}

func DeleteOrderById(ctx *gin.Context) {
	log.Logger().Info("[DeleteOrderById] ", ctx.Request.URL)

}

func GenerateOrder(ctx *gin.Context) {
	log.Logger().Info("[GenerateOrder] ", ctx.Request.URL)
}

func GetOrdersByProductId(ctx *gin.Context) {
	log.Logger().Info("[GetOrdersByProductId] ", ctx.Request.URL)

}

func UpdateOrderInfoById(ctx *gin.Context) {
	log.Logger().Info("[UpdateOrderInfoById] ", ctx.Request.URL)

}
