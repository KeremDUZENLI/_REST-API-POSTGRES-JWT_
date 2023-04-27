package main

import (
	"postgre-project/common/env"
	"postgre-project/controller"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/repository"
	"postgre-project/router"
	"postgre-project/service"
)

func main() {
	env.Load()

	setDatabase()

	router := setAllDependencies()
	router.Run(env.ROUTER)
}

func setDatabase() {
	database.ConnectDB()
	database.ConnectDB().AutoMigrate(&model.Tables{})
}

func setAllDependencies() router.Router {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	controller := controller.NewController(service)
	router := router.NewRouter(controller)

	return router
}
