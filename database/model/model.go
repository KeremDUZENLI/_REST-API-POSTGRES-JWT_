package model

import (
	"gorm.io/gorm"
)

type Tables struct {
	gorm.Model
	FirstName string `gorm:"column:first_name" json:"firstname" validate:"required,min=2,max=100"`
	LastName  string `gorm:"column:last_name" json:"lastname" validate:"required,min=2,max=100"`
	Email     string `gorm:"column:email" json:"email" validate:"email,required"`
	Password  string `gorm:"column:password" json:"password" validate:"required,min=6"`
	UserType  string `gorm:"column:user_type" json:"usertype" validate:"required,eq=ADMIN|eq=USER"`
	Token     string `gorm:"column:token"`
}
