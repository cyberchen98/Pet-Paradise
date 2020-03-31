package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet-paradise/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			utils.Response(ctx, http.StatusUnauthorized, "重新登陆", nil)
			ctx.Abort()
			return
		}
		claims, err := ParseToken(tokenString)
		if err != nil {
			utils.Response(ctx, http.StatusUnauthorized, "权限不足", nil)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserId)
		ctx.Next()
	}
}
