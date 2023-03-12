package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	controller "github.com/nicchunglow/go-fiber-bookstore/controllers"
	"github.com/nicchunglow/go-fiber-bookstore/database"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Post("/users", controller.CreateUser)
	app.Get("/users", controller.GetAllUsers)
	app.Get("/users/:id", controller.GetUser)
	app.Put("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.DeleteUser)
}

func main() {
	app := fiber.New()
	database.ConnectDb()
	SetupRoutes(app)
	port := os.Getenv("PORT")
	fmt.Printf("Server starting at http://localhost:%v", port)
	log.Fatal(app.Listen(":" + port))
}
