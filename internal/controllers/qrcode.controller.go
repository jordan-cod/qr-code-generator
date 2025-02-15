package controllers

import (
	"net/http"

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

	c.Data(http.StatusOK, "image/png", qrCodeImage)
}
