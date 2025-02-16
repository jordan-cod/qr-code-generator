package controllers

import (
	"net/http"
	"qr-code-generator/internal/database"
	"qr-code-generator/internal/models"
	"qr-code-generator/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type QRCodeRequest struct {
	Text string `json:"text" binding:"required"`
}

// GenerateQRCode godoc
// @Summary      Gerar um QR Code
// @Description  Gera um QR Code a partir do texto fornecido no corpo da requisição
// @Accept       json
// @Produce      json
// @Param        request body QRCodeRequest true "Texto a ser convertido em QR Code"
// @Success      200 {file} png "QR Code gerado com sucesso"
// @Failure      400 {object} map[string]string "Requisição inválida"
// @Failure      500 {object} map[string]string "Erro ao gerar o QR Code"
// @Router       /generate [post]
func GenerateQRCode(c *gin.Context) {
	var request QRCodeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Texto inválido"})
		return
	}

	qrCodeImage, err := qrcode.Encode(request.Text, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o QR Code"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	qrcodeRepo := repositories.NewRepository[models.QRCode](database.GetDB())
	qrCodeModel := models.QRCode{
		Text:   request.Text,
		Image:  qrCodeImage,
		UserID: userID.(string),
	}

	if err := qrcodeRepo.Create(c.Request.Context(), &qrCodeModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar o QR Code no banco de dados"})
		return
	}

	c.Data(http.StatusOK, "image/png", qrCodeImage)
}
