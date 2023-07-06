package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
var db *gorm.DB
// InitDB initializes the database connection
func InitDB() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Failed to connect to the database!")
	}

	return db
}
