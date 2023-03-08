package routes

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/nicchunglow/go-fiber-bookstore/database"
	"github.com/nicchunglow/go-fiber-bookstore/models"
)

type User struct {
	//this is the user serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUserMapper(userModel *models.User) User {
	fmt.Print(userModel)
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	DB := database.Database.Db
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	DB.Create(&user)
	responseUser := CreateResponseUserMapper(&models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	return c.Status(200).JSON(responseUser)
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUserMapper(&user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.JSON(responseUsers)
}
func GetUserById(id int) (*User, error) {
	var user models.User
	err := database.Database.Db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	responseUser := CreateResponseUserMapper(&models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	return &responseUser, nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	user, err := GetUserById(id)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.JSON(user)
}
