package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/nicchunglow/dancecircle-backend/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env! \n", err)
	}
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		"test-user",
		os.Getenv("DB_PASSWORD"),
		"testdb",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}
	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{})

	Database = DbInstance{Db: db}

	fmt.Println("Database Connected Successfully")
}
