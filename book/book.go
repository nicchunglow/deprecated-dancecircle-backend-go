package book

import (
	"github.com/gofiber/fiber"
)

func GetAllBooks(c *fiber.Ctx) {
	c.Send("All Books")
}
func GetBook(c *fiber.Ctx) {
	c.Send("All Books")
}
func NewBook(c *fiber.Ctx) {
	c.Send("All Books")
}
func DeleteBook(c *fiber.Ctx) {
	c.Send("All Books")
}
