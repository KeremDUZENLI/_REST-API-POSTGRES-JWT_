package model

import (
	"gorm.io/gorm"
)

type Tables struct {
	gorm.Model
	FirstName string `gorm:"column:first_name" validate:"required,min=2,max=100"`
	LastName  string `gorm:"column:last_name" validate:"required,min=2,max=100"`
	Email     string `gorm:"column:email" validate:"email,required"`
	Password  string `gorm:"column:password" validate:"required,min=6"`
	UserType  string `gorm:"column:user_type" validate:"required,eq=ADMIN|eq=USER"`
	Token     string `gorm:"column:token"`
}
