package routes

import (
	"qr-code-generator/internal/controllers"
	"qr-code-generator/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/register", controllers.SignUp)
			authRoutes.POST("/login", controllers.SignIn)
		}

		qrcodeRoutes := api.Group("/qrcode", middlewares.AuthMiddleware())
		{
			qrcodeRoutes.POST("/generate", controllers.GenerateQRCode)

		}

	}
}
