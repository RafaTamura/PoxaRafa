package controllers

import (
	"log"
	"net/http"
	"os"
	"poxa_rafa/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	expectedUser := os.Getenv("USER")
	expectedPass := os.Getenv("ALO")

	log.Println("USER:", expectedUser)
	log.Println("PASS:", expectedPass)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if body.Username != expectedUser || body.Password != expectedPass {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	token, err := utils.GenerateJWT(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
