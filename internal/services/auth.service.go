package services

import (
	"context"
	"errors"
	"log"
	"qr-code-generator/config"
	"qr-code-generator/internal/database"
	"qr-code-generator/internal/models"
	"qr-code-generator/internal/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(userID string) (string, error) {
	secretKey := config.GetEnv("JWT_SECRET_KEY", "")
	if secretKey == "" {
		log.Panic("Variável de ambiente obrigatória não encontrada: JWT_SECRET_KEY")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func AuthenticateUser(ctx context.Context, email, password string) (string, error) {
	var user *models.User
	userRepo := repositories.NewRepository[models.User](database.GetDB())

	conditions := map[string]interface{}{"email": email}

	user, err := userRepo.GetBy(ctx, conditions)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	err = ComparePassword(user.Password, password)
	if err != nil {
		return "", errors.New("senha incorreta")
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
