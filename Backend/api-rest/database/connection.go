package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {

	var DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", dbHost, dbPort, dbUser, dbPassword, dbName)
	fmt.Println(DSN)

	var error error
	Db, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		panic("Failed to connect to the database!")
	} else {
		log.Println("connect to the database")
	}

}
