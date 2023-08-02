package controller_test

import (
	"testing"

	mocks "github.com/nicchunglow/dancecircle-backend/controller/mock_UserController_test.go" // Import the generated mock package

	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/dancecircle-backend/controller" // Import your controller package
	"github.com/nicchunglow/dancecircle-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	// Create an instance of the MockUserController
	mockUserController := new(mocks.MockUserController)

	// Prepare the Fiber context with a request body
	userRequest := models.User{FirstName: "John", LastName: "Doe"}
	ctx := fiber.New()
	ctx.Request().Body = []byte(`{"first_name":"John","last_name":"Doe"}`)

	// Set up the expectations for the CreateUser method
	mockUserController.On("CreateUser", ctx).Return(nil)

	// Call the actual function that you want to test (in this case, CreateUser)
	err := controller.CreateUser(ctx, mockUserController)

	// Assert that the function behaved as expected
	assert.NoError(t, err)

	// Optionally, you can assert that all the expected calls were made
	mockUserController.AssertExpectations(t)
}

func TestCreateUser_BadRequest(t *testing.T) {
	// Create an instance of the MockUserController
	mockUserController := new(mocks.MockUserController)

	// Prepare the Fiber context with an invalid request body
	ctx := fiber.New()
	ctx.Request().Body = []byte(`invalid json`)

	// Call the actual function that you want to test (in this case, CreateUser)
	err := controller.CreateUser(ctx, mockUserController)

	// Assert that the function returns a 400 error for the bad request
	assert.Equal(t, fiber.StatusBadRequest, ctx.Response().StatusCode)
	assert.Error(t, err)

	// Optionally, you can assert that no unexpected calls were made
	mockUserController.AssertExpectations(t)
}

// Write more test functions for other controller methods using the same approach.
