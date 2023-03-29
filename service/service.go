package service

import (
	"errors"

	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/dto/mapper"
	"postgre-project/middleware"
	"postgre-project/middleware/auth"
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

func FindUser(c *gin.Context, dLI dto.DtoLogIn) (logIn model.Tables, errLogIn error) {
	aMap := mapper.MapperLogIn(&dLI)
	logIn, errLogIn = repository.FindByEmail(aMap)
	if errLogIn != nil {
		return
	}

	if !isEmailValid(aMap.Email, logIn.Email) || !isPasswordValid(logIn.Password, aMap.Password) {
		return model.Tables{}, errors.New("not valid")
	}

	if err := update(logIn); err != nil {
		return
	}

	return
}

func GetUserByID(c *gin.Context, personId int) (model.Tables, error) {
	if err := auth.IsAdmin(c); err != nil {
		return model.Tables{}, err
	}

	return repository.GetInfoByIdFromDatabase(personId)
}

func GetUsersAll(c *gin.Context) ([]model.Tables, error) {
	if err := auth.IsAdmin(c); err != nil {
		return []model.Tables{}, err
	}

	return repository.GetInfosFromDatabase()
}

// ----------------------------------------------------------------
func setValues(person *model.Tables) error {
	person.Password, _ = middleware.HashPassword(person.Password)
	_, errPassword := middleware.HashPassword(person.Password)

	signedToken, errGenerate := token.GenerateToken(person.FirstName, person.LastName, person.Email, person.UserType)
	person.Token = signedToken

	if errPassword != nil && errGenerate != nil {
		return errors.New("error setting values")
	}

	return nil
}

func isUserExist(person model.Tables) bool {
	_, err := repository.FindByEmail(person)

	return err == nil
}

func isEmailValid(mapEmail string, databaseEmail string) bool {
	return mapEmail == databaseEmail
}

func isPasswordValid(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func update(person model.Tables) error {
	signedToken, errGenerate := token.GenerateToken(person.FirstName, person.LastName, person.Email, person.UserType)
	if errGenerate != nil {
		return errors.New("token generation failed")
	}

	if errUpdate := token.UpdateToken(person.ID, signedToken); errUpdate != nil {
		return errors.New("token update failed")
	}

	return nil
}
