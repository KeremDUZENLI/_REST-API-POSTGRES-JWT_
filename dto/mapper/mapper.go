package mapper

import (
	"postgre-project/database/model"
	"postgre-project/dto"
)

// SignUp
func MapperSignUp(d *dto.DtoSignUp) model.Tables {
	return model.Tables{
		Password:  d.Password,
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Email:     d.Email,
		UserType:  d.UserType,
	}
}

// LogIn
func MapperLogIn(d *dto.DtoLogIn) model.Tables {
	return model.Tables{
		Email:    d.Email,
		Password: d.Password,
	}
}

// GetUserById

// GetUsers
