package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	userController "github.com/nicchunglow/dancecircle-backend-go/controllers"
	models "github.com/nicchunglow/dancecircle-backend-go/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll() ([]models.UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]models.UserResponse), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	defer mockRepo.AssertExpectations(t)

	users := []models.UserResponse{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
	}

	mockRepo.On("GetAll").Return(users, nil)

	app := fiber.New()
	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := mockRepo.GetAll()
		if err != nil {
			return err
		}

		return c.JSON(users)
	})

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)

	expectedBodyBytes, _ := json.Marshal(users)
	expectedBody := string(expectedBodyBytes)

	assert.Equal(t, expectedBody, string(body))
}

func TestCreateUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	defer mockRepo.AssertExpectations(t)

	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}

		err = mockRepo.CreateUser(user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		responseUser := userController.CreateResponseUserMapper(user)
		return c.Status(http.StatusOK).JSON(responseUser)
	})

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	defer mockRepo.AssertExpectations(t)

	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}

		err = mockRepo.CreateUser(user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		responseUser := userController.CreateResponseUserMapper(user)
		return c.Status(http.StatusOK).JSON(responseUser)
	})

	user := models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	userJSON, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(userJSON))

	req.Header.Set("Content-Type", "application/json")

	mockRepo.On("CreateUser", user).Return(nil)

	resp, err := app.Test(req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var responseUser models.UserResponse
	err = json.Unmarshal(body, &responseUser)
	assert.Nil(t, err)

	expectedResponseUser := userController.CreateResponseUserMapper(user)
	assert.Equal(t, expectedResponseUser.ID, responseUser.ID)
	assert.Equal(t, expectedResponseUser.FirstName, responseUser.FirstName)
	assert.Equal(t, expectedResponseUser.LastName, responseUser.LastName)
}
