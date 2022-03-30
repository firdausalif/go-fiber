package database

import (
	"fmt"
	"github.com/firdausalif/go-fiber/config"
	"github.com/firdausalif/go-fiber/internal/model"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dbPort := config.Config("DB_PORT")
	port, err := strconv.ParseUint(dbPort, 10, 32)

	if err != nil {
		log.Println("Failed parse por to int")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port,
		config.Config("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")
}
