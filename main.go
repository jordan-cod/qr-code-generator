package main

import (
	"log"
	_ "qr-code-generator/docs"
	"qr-code-generator/internal/config"
	"qr-code-generator/internal/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           Gerator de QR Codes
// @version         1.0
// @description     API para gerar QR Codes

// @host            localhost:8080
// @BasePath        /api
func main() {
	config.LoadEnv()

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router)

	port := config.GetEnv("PORT", "8080")
	log.Printf("Listening on port %s", port)

	router.Run(":" + port)
}
