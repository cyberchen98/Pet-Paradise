package api

import (
	"github.com/gin-gonic/gin"
	"pet-paradise/api/impl"
	"pet-paradise/middleware"
	"pet-paradise/utils"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group("/api/v1")
	devBase := middleware.Cors()
	baseRouter.Use(devBase)

	baseRouter.GET("/ping", LiveProbe)

	baseRouter.POST("/login", impl.Login)
	baseRouter.POST("/register", impl.Register)

	baseRouter.GET("/product/parentProduct", impl.GetParentProduct)
	baseRouter.GET("/product/all/:parentProductName", impl.GetProductInfoByParentName)
	baseRouter.GET("/product", impl.GetProductInfoByName)

	authFunc := middleware.AuthMiddleware()
	authRouter := baseRouter.Group("/user")
	authRouter.Use(authFunc)

	authRouter.GET("/logout", impl.Logout)
	authRouter.GET("/info", impl.GetUserInfo)
	authRouter.PATCH("/info", impl.UpdateUserInfo)
	authRouter.DELETE("/info", impl.DeleteUser)
	authRouter.PATCH("/info/password", impl.UpdateUserPassword)

	authRouter.GET("/address", impl.GetAllAddress)
	authRouter.PUT("/address", impl.AddAddressInfo)
	authRouter.PATCH("/address", impl.UpdateAddressInfo)
	authRouter.DELETE("/address", impl.DeleteAddress)

	authRouter.GET("/order/all", impl.GetAllOrderInfoByUserId)
	authRouter.POST("/order", impl.GenerateOrder)
	authRouter.DELETE("/order", impl.DeleteOrderById)
	authRouter.PATCH("/order", impl.UpdateOrderInfoById)

	adminAuthFunc := middleware.AdminAuthMiddleware()
	adminRouter := baseRouter.Group("/admin")
	adminRouter.Use(authFunc, adminAuthFunc)

	adminRouter.PUT("/product/parentProduct", impl.AddNewParentProduct)
	adminRouter.DELETE("/product/parentProduct", impl.AdminDeleteParentProduct)
	adminRouter.PUT("/product", impl.AddNewProduct)
	adminRouter.PATCH("/product", impl.UpdateProductInfo)
	adminRouter.DELETE("/product", impl.DeleteProduct)
	adminRouter.GET("/order", impl.AdminGetOrdersByProductId)
	adminRouter.PATCH("/order", impl.AdminUpdateOrderInfoById)
	adminRouter.GET("/user", impl.AdminGetUserInfoByName)
	adminRouter.PATCH("/user", impl.AdminUpdateUserRole)

	return r
}

func LiveProbe(ctx *gin.Context) {
	utils.Success(ctx, "pong", nil)
}
