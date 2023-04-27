package controller

import (
	"net/http"
	"postgre-project/dto"
	"postgre-project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user struct {
	service service.PostgreService
}

type User interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
}

func NewController(serv service.PostgreService) User {
	return &user{service: serv}
}

// ----------------------------------------------------------------

func (u user) SignUp(c *gin.Context) {
	var dto dto.DtoSignUp
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return
	}

	if err := u.service.CreateUser(c, dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto)
}

func (u user) LogIn(c *gin.Context) {
	var dto dto.DtoLogIn
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return
	}

	logIn, err := u.service.FindUser(c, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logIn)
}

func (u user) GetUser(c *gin.Context) {
	personId := c.Param("user_id")
	num, _ := strconv.Atoi(personId)

	getUser, err := u.service.GetUserByID(c, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getUser)
}

func (u user) GetUsers(c *gin.Context) {
	usersList, err := u.service.GetUsersAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usersList)
}
