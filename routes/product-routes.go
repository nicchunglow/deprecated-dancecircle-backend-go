package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/nicchunglow/dancecircle-backend/controllers"
)

func ProductRoutes(app *fiber.App) {
	app.Post("/products", controller.CreateProduct)
	app.Get("/products", controller.GetProducts)
	app.Get("/products/:id", controller.GetProduct)
	app.Put("/products/:id", controller.UpdateProduct)
}
