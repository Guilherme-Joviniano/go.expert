package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@gmail.com", "12346578")

	assert.NotEmpty(t, user.Id)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johndoe@gmail.com", user.Email)

	assert.Nil(t, err)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@gmail.com", "12346578")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("12346578"))
	assert.False(t, user.ValidatePassword("any_invalid_password"))
	assert.NotEqual(t, user.Password, "12346578")
}
