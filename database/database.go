package database

import (
	"log"

	"github.com/amirfakhrullah/go-graphql/env"
	"github.com/amirfakhrullah/go-graphql/graph/customTypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(postgres.Open(env.DB_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected Successfully to Database")
	if err := db.AutoMigrate(&customTypes.Todo{}); err != nil {
		return nil, err
	}
	return db, nil
}
