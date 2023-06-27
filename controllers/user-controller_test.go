package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	controller "github.com/nicchunglow/dancecircle-backend/controllers"
	"github.com/nicchunglow/dancecircle-backend/models"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
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

var _ = ginkgo.Describe("UserController", func() {
	var (
		mockUserRepository *MockUserRepository
		app                *fiber.App
	)

	ginkgo.BeforeEach(func() {
		mockUserRepository = &MockUserRepository{}
		userController := controller.UserController(mockUserRepository)
		app = fiber.New()
		app.Get("/users", userController.GetAllUsers)
	})

	ginkgo.It("should return all users", func() {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		resp, err := app.Test(req)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		defer resp.Body.Close()

		gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusOK))

		var responseUsers []controller.User
		err = json.NewDecoder(resp.Body).Decode(&responseUsers)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		gomega.Expect(responseUsers).To(gomega.HaveLen(2))

		gomega.Expect(responseUsers[0].ID).To(gomega.Equal(uint(1)))
		gomega.Expect(responseUsers[0].FirstName).To(gomega.Equal("John"))
		gomega.Expect(responseUsers[0].LastName).To(gomega.Equal("Doe"))

		gomega.Expect(responseUsers[1].ID).To(gomega.Equal(uint(2)))
		gomega.Expect(responseUsers[1].FirstName).To(gomega.Equal("JohnJohn"))
		gomega.Expect(responseUsers[1].LastName).To(gomega.Equal("Doey"))
	})
})
