package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scanner-web-api/auth"
	"scanner-web-api/dbsetup"
	"scanner-web-api/models"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.Usuario
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	record := dbsetup.DB.Where("\"Nombre\" = ? and \"baja\" = 1", request.Username).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "No se encontro un usuario activo con nombre de usuario: " + request.Username})
		context.Abort()
		return
	}

	user.HashPassword(user.Password)
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username, user.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
