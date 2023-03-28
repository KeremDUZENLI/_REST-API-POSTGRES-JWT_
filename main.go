package main

import (
	"fmt"
	"postgre-project/common/env"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Database Open
	env.Load()
	database.ConnectDB()
	database.Instance.AutoMigrate(&model.Tables{})

	// SignUp
	signUp()

	// LogIn
	logIn()

	// GetResult
	getResultById()

	// GetResults
	getResults()

	// Database Close
	database.CloseDB()
}

var c gin.Context
var pl = fmt.Printf

func signUp() {
	service.CreateUser(&c, dto.DtoSignUp{
		Password:     "aaa",
		Token:        "",
		RefreshToken: "",
		FirstName:    "",
		LastName:     "",
		Email:        "aaa",
		UserType:     "",
	})
}

func logIn() {
	find, errFind := service.FindUser(&c, dto.DtoLogIn{
		Email:    "aaaa",
		Password: "aaa",
	})

	pl("\nFinding: \n%v\n", find)
	pl("\nError: \n%v\n", errFind)
}

func getResultById() {
	resById, errId := service.GetUserByID(&c, 1)

	pl("\nResult By ID: \n%v\n", resById)
	pl("\nError Result ID: \n%v\n", errId)
}

func getResults() {
	res, errRes := service.GetUsers(&c)

	for _, v := range res {
		pl("\nResults: \n%v   ---   %v\n", v.Email, v.Password)
	}
	// pl("\nResults: \n%v\n", res)
	pl("\nError Results: \n%v\n", errRes)
}
