package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Name            string `json:"name" validate:"required,min=3,max=50"`
	Email           string `json:"email" validate:"required,email,max=50"`
	Password        string `json:"password" validate:"required,min=8,max=18"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

var users = []SignUpRequest{}

// TODO: Melhorar a validação dos campos para retornar mensagens de erro mais claras
func SignUp(c *gin.Context) {
	var request SignUpRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
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

	users = append(users, request)

	c.JSON(http.StatusOK, gin.H{"message": "Usuário cadastrado com sucesso"})
}

func SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Usuário autenticado com sucesso"})
}
