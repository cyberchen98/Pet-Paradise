package api

import (
	"github.com/gin-gonic/gin"
	"pet-paradise/api/impl"
	"pet-paradise/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group("/api/v1")
	devBase := middleware.Cors()
	baseRouter.Use(devBase)

	baseRouter.POST("/login", impl.Login)
	baseRouter.POST("/register", impl.Register)

	baseRouter.GET("/product/:parentProductName", impl.GetProductInfoByParentName)
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
	authRouter.GET("/order", impl.GetOrderInfoById)
	authRouter.POST("/order", impl.GenerateOrder)
	authRouter.DELETE("/order", impl.DeleteOrderById)
	authRouter.PATCH("/order", impl.UpdateOrderInfoById)

	adminAuthFunc := middleware.AdminAuthMiddleware()
	adminRouter := baseRouter.Group("/admin")
	adminRouter.Use(authFunc, adminAuthFunc)

	adminRouter.POST("/product", impl.AddNewProduct)
	adminRouter.DELETE("/product", impl.DeleteProduct)
	adminRouter.PATCH("/product", impl.UpdateProductInfo)
	adminRouter.PATCH("/order", impl.AdminUpdateOrderInfoById)

	return r
}
