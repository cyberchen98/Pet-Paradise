package middleware

import (
	"github.com/gin-gonic/gin"
	"pet-paradise/utils"
	"strconv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			utils.Fail(ctx, "please login", nil)
			ctx.Abort()
			return
		}
		claims, err := ParseToken(tokenString)
		if err != nil {
			utils.Fail(ctx, "please login again", nil)
			ctx.Abort()
			return
		}

		ctx.Set("uid", strconv.Itoa(claims.UserId))
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if role, ok := ctx.Get("role"); !ok || role != "admin" {
			utils.Fail(ctx, "permission denied", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
