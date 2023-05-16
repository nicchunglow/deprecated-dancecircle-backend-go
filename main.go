package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/dancecircle-backend/database"
	"github.com/nicchunglow/dancecircle-backend/routes"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	routes.UserRoutes(app)
	routes.ProductRoutes(app)
	routes.OrderRoutes(app)
}

func main() {
	app := fiber.New()
	database.ConnectDb()
	SetupRoutes(app)
	port := os.Getenv("PORT")
	fmt.Printf("Server starting at http://localhost:%v", port)
	log.Fatal(app.Listen(":" + port))
}
