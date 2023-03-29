package model

import (
	"gorm.io/gorm"
)

type Tables struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	UserType  string `gorm:"column:user_type"`
	Token     string `gorm:"column:token"`
}
