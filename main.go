package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/nicchunglow/go-fiber-bookstore/database"
	routes "github.com/nicchunglow/go-fiber-bookstore/routes"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	db, err := gorm.Open("mysql", "root:"+dbPassword+"@/fiberBookStore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	database.DBConn = db
	defer database.DBConn.Close()
	fmt.Println("Database Connected Successfully")
}
func main() {
	app := fiber.New()
	InitDatabase()
	app.Get("/", helloWorld)
	routes.SetupRoutes(app)
	port := os.Getenv("PORT")
	fmt.Printf("Server starting at http://localhost:%v", port)
	log.Fatal(app.Listen(":" + port))
}
