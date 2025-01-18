package controllers

import (
	"net/http"
	"time"
	"v1-backend-sebastian-nadal/configs"
	"v1-backend-sebastian-nadal/models"
	"v1-backend-sebastian-nadal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Solicitar recuperación de contraseña
func SolicitarRecuperacion(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var usuario models.Usuario
	if err := configs.DB.
		Preload("Credencial"). // Preload de la relación Credencial
		Where("credenciales.email = ?", request.Email).
		Joins("JOIN credenciales ON credenciales.usuario_id = usuarios.id").
		First(&usuario).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	token := uuid.NewString()
	expiration := time.Now().Add(15 * time.Minute)

	recuperacion := models.Recuperacion{
		Token:     token,
		UsuarioID: usuario.ID,
		ExpiresAt: expiration,
	}
	if err := configs.DB.Create(&recuperacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
		return
	}

	// Este link aparecera en el correo enviado por es sistema
	// Ruta desarrollo local path: '/reset-password/:token'
	link := "https://v1-frontend-sebastian-nadal-production.up.railway.app/reset-password/" + token
	go utils.EnviarEmailRecuperacion(request.Email, link)

	c.JSON(http.StatusOK, gin.H{"message": "Email enviado con las instrucciones. (Aparecerá en tu correo: entre 1 a 10 minutos)"})
}

// Cambiar contraseña
func CambiarContrasena(c *gin.Context) {
	var request struct {
		Token        string `json:"token" binding:"required"`
		NuevaClave   string `json:"nueva_clave" binding:"required,min=8"`
		ConfirmClave string `json:"confirm_clave" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.NuevaClave != request.ConfirmClave {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Las contraseñas no coinciden"})
		return
	}

	var recuperacion models.Recuperacion
	if err := configs.DB.Where("token = ?", request.Token).First(&recuperacion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token inválido"})
		return
	}

	if time.Now().After(recuperacion.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "El token ha expirado"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.NuevaClave), bcrypt.DefaultCost)
	if err := configs.DB.Model(&models.Credencial{}).
		Where("usuario_id = ?", recuperacion.UsuarioID).
		Update("password", string(hashedPassword)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la contraseña"})
		return
	}

	configs.DB.Delete(&recuperacion)
	c.JSON(http.StatusOK, gin.H{"message": "Contraseña actualizada correctamente"})
}
