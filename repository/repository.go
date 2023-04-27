package repository

import (
	"errors"
	"postgre-project/database"
	"postgre-project/database/model"
)

type postgreRepository struct{}

type PostgreRepository interface {
	AddToDatabase(person model.Tables)
	FindByEmail(person model.Tables) (model.Tables, error)
	GetInfoByIdFromDatabase(id int) (model.Tables, error)
	GetInfosFromDatabase() ([]model.Tables, error)
}

func NewRepository() PostgreRepository {
	return &postgreRepository{}
}

// ----------------------------------------------------------------

func (postgreRepository) AddToDatabase(person model.Tables) {
	database.ConnectDB().Table(model.TABLE).Create(&person)
}

func (postgreRepository) FindByEmail(person model.Tables) (model.Tables, error) {
	if err := database.ConnectDB().
		Table(model.TABLE).
		Where("email = ?", person.Email).
		First(&person).Error; err != nil {
		return model.Tables{}, errors.New("email not found")
	}

	return person, nil
}

func (postgreRepository) GetInfoByIdFromDatabase(id int) (model.Tables, error) {
	var person model.Tables
	if err := database.ConnectDB().
		Table(model.TABLE).
		Where("id = ?", id).
		First(&person).
		Error; err != nil {
		return model.Tables{}, errors.New("id not found")
	}

	return person, nil
}

func (postgreRepository) GetInfosFromDatabase() ([]model.Tables, error) {
	var infos []model.Tables
	if err := database.ConnectDB().
		Table(model.TABLE).
		Find(&infos).
		Error; err != nil {
		return []model.Tables{}, errors.New("nothing found")
	}

	return infos, nil
}
