package dto

import "time"

type CrearImagen_prefabricadaRequest struct {
	//Image string `json:"image" binding:"required"`
	//PrefabricadaID uint   `json:"prefabricada_id" binding:"required"`
	Plano bool `form:"plano"`
}

type ActualizarImagen_prefabricadaRequest struct {
	//Image string `json:"image" binding:"required"`
	Plano bool `form:"plano"`
}

type Imagen_prefabricadaResponse struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_At"`
	Image          string    `json:"image"`
	Plano          bool      `json:"plano"`
	PrefabricadaID uint      `json:"prefabricada_id"`
}
