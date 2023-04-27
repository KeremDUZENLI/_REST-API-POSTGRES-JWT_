package database

import (
	"fmt"
	"log"

	"postgre-project/common/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgreDB *gorm.DB

func ConnectDB() *gorm.DB {
	if postgreDB != nil {
		return postgreDB
	}

	url := dbUrl()
	db := gormOpen(url)

	postgreDB = db
	return postgreDB
}

// ----------------------------------------------------------------

func dbUrl() string {
	url_db := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_DBNAME, env.DB_PORT, env.DB_SSL)

	return url_db
}

func gormOpen(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("can not connect to database")
	}

	return db
}
