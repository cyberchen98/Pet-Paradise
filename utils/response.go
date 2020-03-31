package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, msg string, data interface{}) {
	ctx.JSON(httpStatus, gin.H{"code": httpStatus, "msg": msg, "data": data})
}

func Success(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusOK, msg, data)
}

func Fail(ctx *gin.Context, msg string, data interface{}) {
	Response(ctx, http.StatusBadRequest, msg, data)
}
