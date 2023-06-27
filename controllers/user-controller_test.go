package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	controller "github.com/nicchunglow/dancecircle-backend/controllers"
	"github.com/nicchunglow/dancecircle-backend/models"
	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct{}

func (m *MockUserRepository) GetAll() ([]models.User, error) {
	// Mock implementation to return sample users
	mockUsers := []models.User{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "JohnJohn", LastName: "Doey"},
	}
	return mockUsers, nil
}

func TestGetAllUsers(t *testing.T) {
	userController := controller.UserController(&MockUserRepository{})
	app := fiber.New()
	app.Get("/users", userController.GetAllUsers)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseUsers []controller.User
	err = json.NewDecoder(resp.Body).Decode(&responseUsers)
	assert.NoError(t, err)

	assert.Len(t, responseUsers, 2)

	assert.Equal(t, uint(1), responseUsers[0].ID)
	assert.Equal(t, "John", responseUsers[0].FirstName)
	assert.Equal(t, "Doe", responseUsers[0].LastName)

	assert.Equal(t, uint(2), responseUsers[1].ID)
	assert.Equal(t, "JohnJohn", responseUsers[1].FirstName)
	assert.Equal(t, "Doey", responseUsers[1].LastName)
}
