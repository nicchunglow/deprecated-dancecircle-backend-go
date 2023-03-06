package controllers

import "github.com/gofiber/fiber/v2"

func GetAllBooks(c *fiber.Ctx) error {
	return c.SendString("We got all books")
}
