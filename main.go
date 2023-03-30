package main

import (
	"postgre-project/common/env"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/router"
)

func main() {

	// Database Open
	env.Load()
	database.ConnectDB()
	database.Instance.AutoMigrate(&model.Tables{})

	// Router
	router.Router()
}
