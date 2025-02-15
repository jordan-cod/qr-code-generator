package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Middleware Auth")
		c.Next()
	}
}
