package router

import (
	"postgre-project/controller"
	"postgre-project/middleware"

	"github.com/gin-gonic/gin"
)

type router struct {
	controller controller.User
	engine     *gin.Engine
}

type Router interface {
	Run(host string)
}

func NewRouter(cont controller.User) Router {
	return &router{controller: cont}
}

// ----------------------------------------------------------------

func (r router) Run(host string) {
	r.engine = gin.New()

	r.setup()

	r.engine.Run(host)
}

func (r router) setup() {
	r.engine.Use(gin.Logger())

	r.postRoutes()
	r.getRoutes()
}

func (r router) postRoutes() {
	r.engine.POST("/signup", r.controller.SignUp)
	r.engine.POST("/login", r.controller.LogIn)
}

func (r router) getRoutes() {
	r.engine.Use(middleware.Authenticate)
	r.engine.GET("/getuser/:user_id", r.controller.GetUser)
	r.engine.GET("/getusers", r.controller.GetUsers)
}
