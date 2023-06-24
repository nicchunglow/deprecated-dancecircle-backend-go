package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductModel(t *testing.T) {

	user := Product{
		ID:           1,
		CreatedAt:    time.Now(),
		Name:         "John",
		SerialNumber: "123123123",
	}

	// Verify the user properties
	assert.Equal(t, uint(1), user.ID)
	assert.NotZero(t, user.CreatedAt)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "123123123", user.SerialNumber)
}
