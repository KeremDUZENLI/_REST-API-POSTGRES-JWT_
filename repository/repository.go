package repository

import (
	"errors"
	"postgre-project/database"
	"postgre-project/database/model"
)

func AddToDatabase(person model.Tables) {
	database.Instance.Table(model.TABLE).Create(&person)
}

func FindByEmail(person model.Tables) model.Tables {
	database.Instance.
		Table(model.TABLE).
		Where("email = ?", person.Email).
		Find(&person)

	return person
}

func GetInfoByIdFromDatabase(id int) (model.Tables, error) {
	var person model.Tables
	if err := database.Instance.
		Table(model.TABLE).
		Where("id = ?", id).
		First(&person).
		Error; err != nil {
		return model.Tables{}, errors.New("id not found")
	}

	return person, nil
}

func GetInfosFromDatabase() ([]model.Tables, error) {
	var infos []model.Tables
	if err := database.Instance.
		Table(model.TABLE).
		Find(&infos).
		Error; err != nil {
		return []model.Tables{}, errors.New("nothing found")
	}

	return infos, nil
}
