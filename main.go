package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func main() {
	app := fiber.New()
	app.Get("/", helloWorld)
	port := 3000
	fmt.Printf("Server starting at http://localhost:%v", port)
	app.Listen(port)
}
