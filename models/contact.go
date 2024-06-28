package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primaryKey"`
	Email   string `json:"email" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type SwaggerContact struct {
	Id      uint   `json:"id" readOnly:"true"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
