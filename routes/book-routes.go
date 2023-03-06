package bookRoutes

import (
	"github.com/gofiber/fiber"
	"github.com/nicchunglow/go-fiber-bookstore/book"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/book", book.GetAllBooks)
	app.Get("/book/{id}", book.GetBook)
	app.Post("/book", book.NewBook)
	app.Delete("/book", book.DeleteBook)

}
