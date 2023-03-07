package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/go-fiber-bookstore/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/book", controllers.GetAllBooks)
	// app.Get("/book/{id}", book.GetBook)
	// app.Post("/book", book.NewBook)
	// app.Delete("/book", book.DeleteBook)

}
