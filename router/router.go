package router

import (
	"net/http"
	"postgre-project/controller"
	"postgre-project/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	ginRouter := gin.Default()

	ginRouter.POST("/signup", controller.SignUp)
	ginRouter.POST("/login", controller.LogIn)

	ginRouter.Use(middleware.Authenticate)
	ginRouter.GET("/getuser/:user_id", controller.GetUser)
	ginRouter.GET("/getusers", controller.GetUsers)

	http.ListenAndServe(":9999", ginRouter)
}
