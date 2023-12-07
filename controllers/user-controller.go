package controller

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/dancecircle-backend-go/database"
	"github.com/nicchunglow/dancecircle-backend-go/models"
)

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.UserResponse, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
}

type UpdateUserType struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUserMapper(userModel models.User) models.UserResponse {
	return models.UserResponse{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	DB := database.Database.Db
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	DB.Create(&user)
	responseUser := CreateResponseUserMapper(user)
	return c.Status(200).JSON(responseUser)
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []models.UserResponse{}
	for _, user := range users {
		responseUser := CreateResponseUserMapper(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.JSON(responseUsers)
}
func GetUserById(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	var user models.User

	if err := GetUserById(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = GetUserById(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var updateData UpdateUserType

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUserMapper(user)

	return c.Status(http.StatusOK).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = GetUserById(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(204).JSON("Successfully deleted User")
}
