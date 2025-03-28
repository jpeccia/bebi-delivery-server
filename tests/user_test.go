package tests

import (
	"testing"

	"github.com/jpeccia/bebi-delivery-server/internal/database"
	"github.com/jpeccia/bebi-delivery-server/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	database.Connect()

	user := models.User{
		Username: "testuser",
		Name:     "teste",
		Password: "password123",
		Phone:    "123456789",
		Address:  "123 Test ST",
		Role:     "customer",
	}

	result := database.DB.Create(&user)

	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "testuser", user.Username)
}
