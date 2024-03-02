package storage

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func loadDBURI() {
}

func Connect() {
    db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=agoudjil dbname=api port=5017 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    log.Println("Database Connected")
    DB = db
}
