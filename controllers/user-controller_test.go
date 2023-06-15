package controller

import (
	"encoding/json"
	"io/ioutil"
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

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	database.Database.Db = gormDB

	users := []User{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
	}

	columns := []string{"id", "first_name", "last_name"}

	rows := sqlmock.NewRows(columns)
	for _, user := range users {
		rows.AddRow(user.ID, user.FirstName, user.LastName)
	}

	mock.ExpectQuery(`SELECT \* FROM "users"`).WillReturnRows(rows)

	app := fiber.New()
	app.Get("/users", GetAllUsers)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)

	expectedBodyBytes, _ := json.Marshal(users)
	expectedBody := string(expectedBodyBytes)

	assert.Equal(t, expectedBody, string(body))
}
