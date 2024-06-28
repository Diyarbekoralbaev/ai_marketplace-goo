package models

import (
	"gorm.io/gorm"
	"image"
	"time"
)

type AiModels struct {
	gorm.Model
	Name        string    `json:"name" gorm:"unique" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"default:0"`
	Owner       User      `json:"owner" gorm:"foreignKey:ID"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type SwaggerAiModel struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Owner       string  `json:"owner"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type UseModel struct {
	Text  string      `json:"text"`
	Image image.Image `json:"image"`
}
