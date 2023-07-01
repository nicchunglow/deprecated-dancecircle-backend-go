package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	defer mockRepo.AssertExpectations(t)

	users := []User{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
	}

	mockRepo.On("GetAll").Return(users, nil)

	app := fiber.New()
	app.Get("/users", func(c *fiber.Ctx) error {
		// Retrieve users from the mock repository
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
	body, err := ioutil.ReadAll(resp.Body)

	expectedBodyBytes, _ := json.Marshal(users)
	expectedBody := string(expectedBodyBytes)

	assert.Equal(t, expectedBody, string(body))
}
