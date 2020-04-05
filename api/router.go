package api

import (
	"github.com/gin-gonic/gin"
	"pet-paradise/api/impl"
	"pet-paradise/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group("/api/v1")
	baseRouter.POST("/login", impl.Login)
	baseRouter.POST("/register", impl.Register)

	baseRouter.GET("/product/:parentProductName", impl.GetProductInfoByParentName)
	baseRouter.GET("/product", impl.GetProductInfoByName)

	authFunc := middleware.AuthMiddleware()
	authRouter := baseRouter.Group("/")
	authRouter.Use(authFunc)

	authRouter.GET("/logout", impl.Logout)
	authRouter.GET("/info", impl.GetUserInfo)
	authRouter.POST("/info", impl.UpdateUserInfo)
	authRouter.DELETE("/info", impl.DeleteUser)
	authRouter.PATCH("/info", impl.UpdateUserPassword)

	authRouter.GET("/address", impl.GetAllAddress)
	authRouter.POST("/address", impl.AddAddressInfo)
	authRouter.PATCH("/address", impl.UpdateAddressInfo)
	authRouter.DELETE("/address", impl.DeleteAddress)

	authRouter.GET("/order", impl.GetOrderInfoByUserId)
	authRouter.POST("/order", impl.GenerateOrder)
	authRouter.DELETE("/order", impl.DeleteOrderById)

	adminAuthFunc := middleware.AdminAuthMiddleware()
	adminRouter := authRouter.Group("/admin")
	adminRouter.Use(adminAuthFunc)

	adminRouter.POST("/product", impl.AddNewProduct)
	adminRouter.DELETE("/product", impl.DeleteProduct)
	adminRouter.PUT("/product", impl.UpdateProductInfo)
	adminRouter.PATCH("/order", impl.UpdateOrderInfoById)
	adminRouter.GET("/order", impl.GetOrdersByProductId)

	return r
}
