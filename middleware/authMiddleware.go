package middleware

import (
	"net/http"
	"postgre-project/database/model"
	"postgre-project/middleware/token"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("client_token")
	if !tokenIsExist(c, clientToken) || !tokenIsValid(c, clientToken) {
		return
	}

	c.Next()
}

func tokenIsExist(c *gin.Context, token string) bool {
	if token == model.NONE {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token does not exist"})
		c.Abort()
		return false
	}

	return true
}

func tokenIsValid(c *gin.Context, givingToken string) bool {
	_, err := token.ValidateToken(givingToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token is not valid"})
		c.Abort()
		return false
	}

	return true
}
