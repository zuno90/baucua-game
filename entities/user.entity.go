package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

type UserInput struct {
	Username string `json:"username" validate:"required,min=6,max=20"`
	Password string `json:"password" validate:"required"`
}
