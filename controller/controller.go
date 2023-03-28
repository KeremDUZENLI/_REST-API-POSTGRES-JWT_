package controller

import (
	"net/http"
	"postgre-project/dto"
	"postgre-project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var dto dto.DtoSignUp
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return
	}

	if err := service.CreateUser(c, dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto)
}

func LogIn(c *gin.Context) {
	var dto dto.DtoLogIn
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return
	}

	logIn, err := service.FindUser(c, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logIn)
}

func GetUser(c *gin.Context) {
	personId := c.Param("user_id")
	num, _ := strconv.Atoi(personId)

	getUser, err := service.GetUserByID(c, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, getUser)
}

func GetUsers(c *gin.Context) {
	usersList, err := service.GetUsersAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usersList)
}
