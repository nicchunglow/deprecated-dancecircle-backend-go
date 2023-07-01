package models

import (
	"encoding/json"
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

func TestUserResponse(t *testing.T) {
	// Create a test user
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	// Serialize the user to JSON
	jsonData, err := json.Marshal(user)
	assert.Nil(t, err)

	// Deserialize the JSON back to a user object
	var deserializedUser User
	err = json.Unmarshal(jsonData, &deserializedUser)
	assert.Nil(t, err)

	// Assert that the deserialized user matches the original user
	assert.Equal(t, user.ID, deserializedUser.ID)
	assert.Equal(t, user.FirstName, deserializedUser.FirstName)
	assert.Equal(t, user.LastName, deserializedUser.LastName)
}
