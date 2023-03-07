package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/go-fiber-bookstore/database"
	routes "github.com/nicchunglow/go-fiber-bookstore/routes"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func main() {
	app := fiber.New()
	database.ConnectDb()
	app.Get("/", helloWorld)
	routes.SetupRoutes(app)
	port := os.Getenv("PORT")
	fmt.Printf("Server starting at http://localhost:%v", port)
	log.Fatal(app.Listen(":" + port))
}
