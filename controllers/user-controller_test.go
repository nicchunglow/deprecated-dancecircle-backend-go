package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/nicchunglow/dancecircle-backend/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Set the mock GORM DB instance in the database package
	// Convert the SQL database connection to a GORM DB instance
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	// Set the GORM DB instance in the database package
	database.Database.Db = gormDB

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
		AddRow(1, "John", "Doe").
		AddRow(2, "Jane", "Smith")

	mock.ExpectQuery("SELECT * FROM users").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/users", GetAllUsers)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Add assertions for the response body if necessary
}
