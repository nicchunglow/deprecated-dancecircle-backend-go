package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestOrderRoutes(t *testing.T) {
	// Create a new Fiber app
	app := fiber.New()

	// Register the order routes
	UserRoutes(app)

	t.Run("POST /users", func(t *testing.T) {
		// Create a request with a dummy payload
		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		resp, err := app.Test(req)

		// Assert the response status code and body
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Add more assertions for the response body if necessary
	})
}