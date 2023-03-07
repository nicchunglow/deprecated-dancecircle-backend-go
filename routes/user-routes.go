package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/go-fiber-bookstore/database"
	"github.com/nicchunglow/go-fiber-bookstore/models"
)

type User struct {
	//this is the user serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	DB := database.Database.Db
	var user models.User
	DB.Create(&user)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func GetAllUsers(c *fiber.Ctx) error {
	DB := database.Database.Db
	users := []models.User{}
	DB.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}
