package config

import (
	"fmt"
	"log"

	"os"

	"aymane/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println(db)
	}

	DB = db

	db.AutoMigrate(&model.User{})
}
