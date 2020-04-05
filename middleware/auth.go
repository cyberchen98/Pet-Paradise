package middleware

import (
	"github.com/gin-gonic/gin"
	"pet-paradise/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			utils.Fail(ctx, "请先登陆", nil)
			ctx.Abort()
			return
		}
		claims, err := ParseToken(tokenString)
		if err != nil {
			utils.Fail(ctx, "重新登陆", nil)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserId)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if role, ok :=ctx.Get("role"); !ok || role!="admin" {
			utils.Fail(ctx, "权限不足", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
