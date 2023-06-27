package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/dancecircle-backend/database"
	"github.com/nicchunglow/dancecircle-backend/models"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserRepository interface {
	GetAll() ([]models.User, error)
}

type UserControllerType struct {
	UserRepo UserRepository
}

func CreateResponseUserMapper(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func UserController(userRepo UserRepository) *UserControllerType {
	return &UserControllerType{
		UserRepo: userRepo,
	}
}

func (c *UserControllerType) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := c.UserRepo.GetAll()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	responseUsers := make([]User, 0, len(users))
	for _, user := range users {
		responseUser := CreateResponseUserMapper(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return ctx.JSON(responseUsers)
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

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUserMapper(user)

	return c.Status(200).JSON(responseUser)
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
