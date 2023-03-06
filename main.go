package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/nicchunglow/go-fiber-bookstore/database"
	bookRoutes "github.com/nicchunglow/go-fiber-bookstore/routes"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
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
	bookRoutes.SetupRoutes(app)
	port := os.Getenv("PORT")
	fmt.Printf("Server starting at http://localhost:%v", port)
	app.Listen(port)
}
