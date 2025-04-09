package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
