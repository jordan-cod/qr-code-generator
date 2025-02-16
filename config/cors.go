package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(router *gin.Engine) {
	allowedOrigins := GetEnv("CORS_ALLOWED_ORIGINS", "*")
	allowedMethods := GetEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE")
	allowedHeaders := GetEnv("CORS_ALLOWED_HEADERS", "Origin,Content-Type,Authorization")
	allowCredentials := GetEnv("CORS_ALLOW_CREDENTIALS", "true") == "true"

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowMethods:     []string{allowedMethods},
		AllowHeaders:     []string{allowedHeaders},
		AllowCredentials: allowCredentials,
	}))
}
