package database

import (
	"fmt"
	"log"

	"rakaminbtpn/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "root"
	password = "root"
	dbPort   = "3307"
	dbname   = "rakaminbtpn"
	db       *gorm.DB
	err      error
)

func Start_db() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	fmt.Println("Connection success to database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func Get_db() *gorm.DB {
	return db
}
