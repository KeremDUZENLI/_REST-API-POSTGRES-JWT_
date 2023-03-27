package main

import (
	"fmt"
	"postgre-project/common/env"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/repository"
)

func main() {

	// Database
	env.Load()
	database.ConnectDB()
	database.Instance.AutoMigrate(&model.Tables{})

	// SignUp
	repository.AddToDatabase(dto.DtoSignUp{
		Password:     "aaa",
		Token:        "",
		RefreshToken: "",
		FirstName:    "",
		LastName:     "",
		Email:        "aaa",
		UserType:     "",
	})

	// LogIn
	find, err := repository.FindInDatabase(dto.DtoLogIn{
		Email:    "aaa",
		Password: "aaa",
	})

	// RESULTS
	pl := fmt.Printf
	pl("\nFinding: \n%v\n", find.CreatedAt)
	pl("\nError: \n%v\n", err)

	database.CloseDB()
}
