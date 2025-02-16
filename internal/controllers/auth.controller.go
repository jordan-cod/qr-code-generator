package controllers

import (
	"net/http"
	"qr-code-generator/internal/database"
	"qr-code-generator/internal/models"
	"qr-code-generator/internal/repositories"
	"qr-code-generator/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Name            string `json:"name" validate:"required,min=3,max=50"`
	Email           string `json:"email" validate:"required,email,max=50"`
	Password        string `json:"password" validate:"required,min=8,max=18"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

// TODO: Melhorar a validação dos campos para retornar mensagens de erro mais claras
func Register(c *gin.Context) {
	var userRepo = repositories.NewRepository[models.User](database.GetDB())
	var request SignUpRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = userRepo.Create(c.Request.Context(), &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar o usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário cadastrado com sucesso"})
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=8,max=18"`
}

func Login(c *gin.Context) {
	var request SignInRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	token, err := services.AuthenticateUser(c.Request.Context(), request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário autenticado com sucesso", "token": token})
}
