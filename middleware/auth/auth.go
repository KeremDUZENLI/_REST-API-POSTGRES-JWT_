package auth

import (
	"errors"
	"postgre-project/database/model"

	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) error {
	userType := c.Request.Header.Get("user_type")

	if userType != model.ADMIN {
		return errors.New("usertype is not ADMIN")
	}

	return nil
}
