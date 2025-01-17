package models

import "time"

type Adicional struct {
	ID              uint       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt       time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	NombreAdicional string     `gorm:"column:nombre_adicional" json:"nombre_adicional"`
	ValorAdicional  float64    `gorm:"column:valor_adicional" json:"valor_adicional"` // Usamos float64 para soportar decimales
	PrecioID        uint       `gorm:"column:precio_id" json:"precio_id"`
	Precio          Precio     `gorm:"foreignKey:PrecioID"`
}

func (Adicional) TableName() string {
	return "adicionales"
}
