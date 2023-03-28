package repository

import (
	"errors"
	"postgre-project/database"
	"postgre-project/database/model"
	"postgre-project/dto"
	"postgre-project/dto/mapper"
	"postgre-project/middleware"
	"postgre-project/middleware/token"

	"golang.org/x/crypto/bcrypt"
)

func AddToDatabase(dSU dto.DtoSignUp) {
	aMap := mapper.MapperSignUp(&dSU)

	if isUserExist(aMap) {
		return
	}

	setValues(&aMap)

	database.Instance.Table(model.TABLE).Create(&aMap)
}

func FindInDatabase(dLI dto.DtoLogIn) (model.Tables, error) {
	aMap := mapper.MapperLogIn(&dLI)

	if !isUserExist(aMap) || !isPasswordExist(aMap, aMap.Password) {
		return model.Tables{}, errors.New("not valid")
	}

	return findByEmail(aMap)
}

func GetInfoByIdFromDatabase(id int) (model.Tables, error) {
	var person model.Tables
	if err := database.Instance.Table(model.TABLE).Where("id = ?", id).First(&person).Error; err != nil {
		return model.Tables{}, errors.New("id not found")
	}

	return person, nil
}

func GetInfosFromDatabase() ([]model.Tables, error) {
	var infos []model.Tables
	if err := database.Instance.Table(model.TABLE).Find(&infos).Error; err != nil {
		return []model.Tables{}, errors.New("nothing found")
	}

	return infos, nil
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
	_, err := findByEmail(person)

	return err == nil
}

func isPasswordExist(person model.Tables, password string) bool {
	res, _ := findByEmail(person)
	err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))

	return err == nil
}

func findByEmail(person model.Tables) (model.Tables, error) {
	return person, database.
		Instance.
		Table(model.TABLE).
		Where("email = ?", person.Email).
		First(&person).
		Error
}
