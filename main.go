package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/go-fiber-bookstore/database"
	"github.com/nicchunglow/go-fiber-bookstore/routes"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/users", routes.GetAllUsers)
	app.Post("/users", routes.CreateUser)
	// app.Get("/book", controllers.GetAllBooks)
	// app.Get("/book/{id}", book.GetBook)
	// app.Post("/book", book.NewBook)
	// app.Delete("/book", book.DeleteBook)

}

func main() {
	app := fiber.New()
	database.ConnectDb()
	SetupRoutes(app)
	port := os.Getenv("PORT")
	fmt.Printf("Server starting at http://localhost:%v", port)
	log.Fatal(app.Listen(":" + port))
}
