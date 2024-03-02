package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadDBURI() {
}

func Connect() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Database Connected")
	DB = db
}
