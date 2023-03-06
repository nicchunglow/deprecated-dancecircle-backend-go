package routes

import (
	"github.com/gofiber/fiber/v2"
)

func getAllBooks(c *fiber.Ctx) error {
	return c.SendString("We got all books")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/book", getAllBooks)
	// app.Get("/book/{id}", book.GetBook)
	// app.Post("/book", book.NewBook)
	// app.Delete("/book", book.DeleteBook)

}
