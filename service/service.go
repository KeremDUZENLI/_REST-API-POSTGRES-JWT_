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
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type postgreService struct {
	repository repository.PostgreRepository
}

type PostgreService interface {
	CreateUser(c *gin.Context, dSU dto.DtoSignUp) error
	FindUser(c *gin.Context, dLI dto.DtoLogIn) (logIn model.Tables, errLogIn error)
	GetUserByID(c *gin.Context, personId int) (model.Tables, error)
	GetUsersAll(c *gin.Context) ([]model.Tables, error)
}

func NewService(repo repository.PostgreRepository) PostgreService {
	return &postgreService{repository: repo}
}

// ----------------------------------------------------------------

func (pS postgreService) CreateUser(c *gin.Context, dSU dto.DtoSignUp) error {
	aMap := mapper.MapperSignUp(&dSU)
	if pS.isUserExist(aMap) || !isObeyRules(aMap) {
		return errors.New("user already exists or can not created because out of rules")
	}

	setValues(&aMap)
	pS.repository.AddToDatabase(aMap)

	return nil
}

func (pS postgreService) FindUser(c *gin.Context, dLI dto.DtoLogIn) (logIn model.Tables, errLogIn error) {
	aMap := mapper.MapperLogIn(&dLI)
	logIn, errLogIn = pS.repository.FindByEmail(aMap)
	if errLogIn != nil {
		return
	}

	if !isEmailValid(aMap.Email, logIn.Email) || !isPasswordValid(logIn.Password, aMap.Password) {
		return model.Tables{}, errors.New("email or password is not valid")
	}

	if err := updateTokenUpdatedat(logIn); err != nil {
		return
	}

	return pS.repository.GetInfoByIdFromDatabase(int(logIn.ID))
}

func (pS postgreService) GetUserByID(c *gin.Context, personId int) (model.Tables, error) {
	if err := auth.IsAdmin(c); err != nil {
		return model.Tables{}, err
	}

	return pS.repository.GetInfoByIdFromDatabase(personId)
}

func (pS postgreService) GetUsersAll(c *gin.Context) ([]model.Tables, error) {
	if err := auth.IsAdmin(c); err != nil {
		return []model.Tables{}, err
	}

	return pS.repository.GetInfosFromDatabase()
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

func (pS postgreService) isUserExist(person model.Tables) bool {
	_, err := pS.repository.FindByEmail(person)

	return err == nil
}

func isObeyRules(person model.Tables) bool {
	return validator.New().Struct(person) == nil
}

func isEmailValid(mapEmail string, databaseEmail string) bool {
	return mapEmail == databaseEmail
}

func isPasswordValid(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func updateTokenUpdatedat(person model.Tables) error {
	signedToken, errGenerate := token.GenerateToken(person.FirstName, person.LastName, person.Email, person.UserType)
	if errGenerate != nil {
		return errors.New("token generation failed")
	}

	if errUpdate := token.UpdateToken(person.ID, signedToken); errUpdate != nil {
		return errors.New("token update failed")
	}

	return nil
}
