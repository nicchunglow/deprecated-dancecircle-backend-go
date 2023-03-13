package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/nicchunglow/go-fiber-bookstore/controllers"
)

func OrderRoutes(app *fiber.App) {
	app.Post("/orders", controller.CreateOrder)
	app.Get("/orders", controller.GetOrders)
	app.Get("/orders/:id", controller.GetOrder)
}
