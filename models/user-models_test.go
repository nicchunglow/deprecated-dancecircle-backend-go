package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	// Create a sample user
	user := User{
		ID:        1,
		CreatedAt: time.Now(),
		FirstName: "John",
		LastName:  "Doe",
	}

	// Verify the user properties
	assert.Equal(t, uint(1), user.ID)
	assert.NotZero(t, user.CreatedAt)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
}
