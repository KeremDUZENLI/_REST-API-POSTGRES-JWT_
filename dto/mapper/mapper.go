package mapper

import (
	"postgre-project/database/model"
	"postgre-project/dto"
)

func MapperSignUp(d *dto.DtoSignUp) model.Tables {
	return model.Tables{
		Password:     d.Password,
		Token:        d.Token,
		RefreshToken: d.RefreshToken,

		FirstName: d.FirstName,
		LastName:  d.LastName,
		Email:     d.Email,
		UserType:  d.UserType,
	}
}

func MapperLogIn(d *dto.DtoLogIn) model.Tables {
	return model.Tables{
		Email:    d.Email,
		Password: d.Password,
	}
}
