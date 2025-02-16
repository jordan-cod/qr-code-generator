package services

import (
	"context"
	"errors"
	"qr-code-generator/internal/database"
	"qr-code-generator/internal/models"
	"qr-code-generator/internal/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(userID string) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    userID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
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
