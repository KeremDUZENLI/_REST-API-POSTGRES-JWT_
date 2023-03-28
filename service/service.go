package service

import (
	"errors"

	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/dto/mapper"
	"postgre-project/middleware"
	"postgre-project/middleware/token"
	"postgre-project/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context, dSU dto.DtoSignUp) error {
	aMap := mapper.MapperSignUp(&dSU)

	if isUserExist(aMap) {
		return errors.New("user already exists")
	}

	setValues(&aMap)
	repository.AddToDatabase(aMap)
	return nil
}

func FindUser(c *gin.Context, dLI dto.DtoLogIn) (model.Tables, error) {
	aMap := mapper.MapperLogIn(&dLI)

	if !isUserExist(aMap) || !isPasswordExist(aMap, aMap.Password) {
		return model.Tables{}, errors.New("not valid")
	}

	return repository.FindByEmail(aMap)
}

func GetUserByID(c *gin.Context, personId int) (model.Tables, error) {
	return repository.GetInfoByIdFromDatabase(personId)
}

func GetUsersAll(c *gin.Context) ([]model.Tables, error) {
	return repository.GetInfosFromDatabase()
}

// service
func setValues(person *model.Tables) error {
	person.Password, _ = middleware.HashPassword(person.Password)
	_, errPassword := middleware.HashPassword(person.Password)

	token, refreshToken, errToken := token.GenerateToken(person.FirstName, person.LastName, person.Email, person.UserType)
	person.Token = token
	person.RefreshToken = refreshToken

	if errPassword != nil && errToken != nil {
		return errors.New("error setValues")
	}

	return nil
}

func isUserExist(person model.Tables) bool {
	_, err := repository.FindByEmail(person)

	return err == nil
}

func isPasswordExist(person model.Tables, password string) bool {
	finding, _ := repository.FindByEmail(person)
	err := bcrypt.CompareHashAndPassword([]byte(finding.Password), []byte(password))

	return err == nil
}
