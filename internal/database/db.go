package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	dsn := "host=localhost user=postgres password=postgres dbname=Register_school port=5432 sslmode=disable TimeZone=GMT-4"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{log.Fatal("Error connecting to database", err)}
	log.Println("Connected")
}