package middlewares

import (
	"errors"
	"log"
	"net/http"
	"qr-code-generator/config"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func getSecretKey() []byte {
	SecretKey := config.GetEnv("JWT_SECRET_KEY", "")
	if SecretKey == "" {
		log.Panic("Variável de ambiente obrigatória não encontrada: JWT_SECRET_KEY")
	}
	return []byte(SecretKey)
}

func validateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return getSecretKey(), nil
	})

	if err != nil || !token.Valid {
		return nil, nil, errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errors.New("não foi possível ler as claims do token")
	}

	return token, claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]
		_, claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido: user_id ausente"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		log.Printf("Usuário autenticado com sucesso: %s", userID)

		c.Next()
	}
}
