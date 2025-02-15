package routes

import (
	"qr-code-generator/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/generate", controllers.GenerateQRCode)
	}
}
