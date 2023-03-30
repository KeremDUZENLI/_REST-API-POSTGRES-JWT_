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

	database.ConnectDB()
	database.Instance.AutoMigrate(&model.Tables{})

	router := setAllDependencies()
	router.Run(env.ROUTER)
}

func setAllDependencies() router.Router {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	controller := controller.NewController(service)
	router := router.NewRouter(controller)

	return router
}
