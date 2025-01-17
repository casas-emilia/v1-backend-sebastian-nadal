package dto

import "time"

type CrearAdicionalRequest struct {
	NombreAdicional string  `json:"nombre_adicional" binding:"required"`
	ValorAdicional  float64 `json:"valor_adicional" binding:"required"`
	//PrecioID      uint   `json:"precio_id" binding:"required"`
}

type ActualizarAdicionalRequest struct {
	NombreAdicional string  `json:"nombre_adicional" binding:"required"`
	ValorAdicional  float64 `json:"valor_adicional" binding:"required"`
}

type AdicionalResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	NombreAdicional string    `json:"nombre_adicional" binding:"required"`
	ValorAdicional  float64   `json:"valor_adicional" binding:"required"`
	PrecioID        uint      `json:"precio_id"`
}
