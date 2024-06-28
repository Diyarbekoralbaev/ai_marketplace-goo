package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primaryKey"`
	Username  string  `json:"username" gorm:"unique" binding:"required"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email" gorm:"unique"`
	Password  string  `json:"password"`
	Balance   float64 `json:"balance" gorm:"default:0"`
}

type SwaggerUser struct {
	Username  string  `json:"username"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Balance   float64 `json:"balance"`
	Password  string  `json:"password" writeOnly:"true"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
