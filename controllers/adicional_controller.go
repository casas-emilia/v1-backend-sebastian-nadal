package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"v1-backend-sebastian-nadal/configs"
	"v1-backend-sebastian-nadal/dto"
	"v1-backend-sebastian-nadal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Función para Crear un Adicional
func CrearAdicional(c *gin.Context) {
	var request dto.CrearAdicionalRequest
	var adicional models.Adicional
	var adicionalResponse dto.AdicionalResponse

	idParamPrecio := c.Param("precioID")
	precioID, err := strconv.ParseUint(idParamPrecio, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Precio inválido")
		return
	}

	// Validamos el body
	if err := c.ShouldBindJSON(&request); err != nil {
		HandleError(c, nil, http.StatusBadRequest, "Error de datos")
		return
	}

	// Creamos el adicional
	adicional.NombreAdicional = request.NombreAdicional
	adicional.ValorAdicional = request.ValorAdicional
	adicional.PrecioID = uint(precioID)

	// Guardamos en la base de datos
	if err := configs.DB.Save(&adicional).Error; err != nil {
		HandleError(c, err, http.StatusInternalServerError, "Error, no se pudo guardar el Adicional")
		return
	}

	adicionalResponse = dto.AdicionalResponse{
		ID:              adicional.ID,
		CreatedAt:       adicional.CreatedAt,
		UpdatedAt:       adicional.UpdatedAt,
		NombreAdicional: adicional.NombreAdicional,
		ValorAdicional:  adicional.ValorAdicional,
		PrecioID:        adicional.PrecioID,
	}

	// Respuesta
	c.JSON(http.StatusCreated, gin.H{
		"mesaage":     "Adicional creado exitosamente",
		"Adicionales": adicionalResponse,
	})
}

// Función para obtener todos los Adicionales
func ObtenerAdicionales(c *gin.Context) {
	var adicionales []models.Adicional
	var adicionalesResponse []dto.AdicionalResponse

	idParamPrecio := c.Param("precioID")
	precioID, err := strconv.ParseUint(idParamPrecio, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Precio inválido")
		return
	}

	// Buscar todos los Adicionales de un precio
	if err := configs.DB.Where("precio_id = ?", precioID).Where("deleted_at IS NULL").Find(&adicionales).Error; err != nil {
		HandleError(c, err, http.StatusInternalServerError, "Error al obtener Incluyes")
		return
	}

	// Verificar si no se encontraron registros
	if len(adicionales) == 0 {
		HandleError(c, nil, http.StatusNotFound, "Sin Adicionales")
		return
	}

	for _, adicional := range adicionales {
		adicionalesResponse = append(adicionalesResponse, dto.AdicionalResponse{
			ID:              adicional.ID,
			CreatedAt:       adicional.CreatedAt,
			UpdatedAt:       adicional.UpdatedAt,
			NombreAdicional: adicional.NombreAdicional,
			ValorAdicional:  adicional.ValorAdicional,
			PrecioID:        adicional.PrecioID,
		})
	}

	// Mostrar/enviar Adicional
	c.JSON(http.StatusOK, gin.H{
		"Adicionales": adicionalesResponse,
	})
}

// Función para obtener un Adicinoal de acuerdo a su ID
func ObtenerAdicional(c *gin.Context) {
	var adicional models.Adicional
	var adicionalResponse dto.AdicionalResponse

	idParamPrecio := c.Param("precioID")
	precioID, err := strconv.ParseUint(idParamPrecio, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Precio inválido")
		return
	}

	idParamIncluye := c.Param("adicionalID")
	adicionalID, err := strconv.ParseUint(idParamIncluye, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Adicional inválido")
		return
	}

	// Buscar Adicional en la base de datos de acuerdo a su ID y al Precio que pertenece
	if err := configs.DB.Where("precio_id = ?", precioID).Where("deleted_at IS NULL").First(&adicional, adicionalID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			HandleError(c, nil, http.StatusNotFound, "Sin Adicionales")
			return
		}
		HandleError(c, err, http.StatusInternalServerError, "Error al obtener Adicional")
		return
	}

	adicionalResponse = dto.AdicionalResponse{
		ID:              adicional.ID,
		CreatedAt:       adicional.CreatedAt,
		UpdatedAt:       adicional.UpdatedAt,
		NombreAdicional: adicional.NombreAdicional,
		ValorAdicional:  adicional.ValorAdicional,
		PrecioID:        adicional.PrecioID,
	}

	// Mostrar/enviar Adicional
	c.JSON(http.StatusBadRequest, gin.H{
		"Adicionales": adicionalResponse,
	})
}

// Función para Actualizar datos de un Adicional
func ActualizarAdicional(c *gin.Context) {
	var request dto.ActualizarAdicionalRequest
	var adicional models.Adicional
	var adicionalResponse dto.AdicionalResponse

	idParamPrecio := c.Param("precioID")
	precioID, err := strconv.ParseUint(idParamPrecio, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Precio inválido")
		return
	}

	idParamIncluye := c.Param("adicionalID")
	adicionalID, err := strconv.ParseUint(idParamIncluye, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Adicional inválido")
		return
	}

	// Bind JSON del request a la estructura del dto
	if err := c.ShouldBindJSON(&request); err != nil {
		HandleError(c, nil, http.StatusBadRequest, "Error de datos"+err.Error())
		return
	}

	// Buscamos el Adicional en la base da datos
	if err := configs.DB.Where("precio_id", precioID).Where("deleted_at IS NULL").First(&adicional, adicionalID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			HandleError(c, nil, http.StatusNotFound, "Sin Adicionales")
			return
		}
		HandleError(c, err, http.StatusInternalServerError, "Error al obtener Adicional")
		return
	}

	// Actualizar datos del Adicional
	adicional.NombreAdicional = request.NombreAdicional
	adicional.ValorAdicional = request.ValorAdicional

	// Guardar los cambios en la base de datos
	if err := configs.DB.Save(&adicional).Error; err != nil {
		HandleError(c, err, http.StatusBadRequest, "Error, no se pudo actualizar los datos del Adicional")
		return
	}

	adicionalResponse = dto.AdicionalResponse{
		ID:              adicional.ID,
		CreatedAt:       adicional.CreatedAt,
		UpdatedAt:       adicional.UpdatedAt,
		NombreAdicional: adicional.NombreAdicional,
		ValorAdicional:  adicional.ValorAdicional,
		PrecioID:        adicional.PrecioID,
	}

	// Mostrar/enviar mensaje de éxito y Adicional
	c.JSON(http.StatusOK, gin.H{
		"message":     "Datos actualizados con éxito",
		"Adicionales": adicionalResponse,
	})
}

// Función para eliminar logicamente un adicional
func EliminarAdicional(c *gin.Context) {
	var adicional models.Adicional

	idParamPrecio := c.Param("precioID")
	precioID, err := strconv.ParseUint(idParamPrecio, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Precio inválido")
		return
	}

	idParamIncluye := c.Param("adicionalID")
	adicionalID, err := strconv.ParseUint(idParamIncluye, 10, 64)
	if err != nil {
		HandleError(c, nil, http.StatusBadRequest, "ID Adicional inválido")
		return
	}

	// Buscar Adicional en la Base de datos
	if err := configs.DB.Where("precio_id = ?", precioID).Where("deleted_at IS NULL").First(&adicional, adicionalID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			HandleError(c, nil, http.StatusNotFound, "Sin Adicionales")
			return
		}
		HandleError(c, err, http.StatusInternalServerError, "Error al obtener Adicional")
		return
	}

	// Verificamos si el Adicional ya se encuentra eliminado
	if adicional.DeletedAt != nil && !adicional.DeletedAt.IsZero() {
		HandleError(c, nil, http.StatusBadRequest, "El Adicional ya se encuentra eliminado")
		return
	}

	// Poner fecha y hora de la eliminación lógica
	now := time.Now()
	adicional.DeletedAt = &now

	// Guardar eliminación lógica en la base de datos
	if err := configs.DB.Save(&adicional).Error; err != nil {
		HandleError(c, err, http.StatusInternalServerError, "Error, no se pudo eliminar el Adicional")
		return
	}

	// Mostrar/enviar mensaje de eliminación lógica exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Adicional eliminado exitosamente",
	})
}
