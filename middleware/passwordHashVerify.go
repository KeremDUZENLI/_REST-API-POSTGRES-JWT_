package middleware

import (
	"postgre-project/database/model"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	encryptionSize := 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), encryptionSize)
	if err != nil {
		return model.NONE, err
	}
	return string(bytes), nil
}

func VerifyPassword(password string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(password))

	return err == nil
}
