package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/nicchunglow/dancecircle-backend/controllers"
)

func UserRoutes(app *fiber.App) {
	app.Post("/users", controller.CreateUser)
	app.Get("/users", controller.GetAllUsers)
	app.Get("/users/:id", controller.GetUser)
	app.Put("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.DeleteUser)
}
