package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	dbPassword := os.Getenv("DB_PASSWORD")
	db, err := gorm.Open("mysql", "root:"+dbPassword+"@/fiberBookStore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}
	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{Db: db}

	fmt.Println("Database Connected Successfully")
}
