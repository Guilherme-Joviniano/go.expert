package database

import (
	"testing"

	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewUserRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	repository := NewUserRepository(db)
	assert.NotNil(t, repository)
}

func TestUserRepository_FindByEmailWhenEmailIsNotValid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	repository := NewUserRepository(db)

	user, err := entity.NewUser("any_name", "any_valid_email@gmail.com", "any_password")

	if err != nil {
		t.Error(err)
	}

	db.Create(user)

	user, err = repository.FindByEmail("invalid_email@gmail.com")

	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "record not found")

}

func TestUserRepository_FindByEmailWhenEmailIsValid(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	repository := NewUserRepository(db)

	user, err := entity.NewUser("any_name", "any_valid_email@gmail.com", "any_password")

	if err != nil {
		t.Error(err)
	}

	db.Create(user)

	userFound, err := repository.FindByEmail("any_valid_email@gmail.com")

	assert.NotNil(t, userFound)
	assert.Equal(t, user.Id, userFound.Id)
	assert.Equal(t, user.Password, userFound.Password)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)

	assert.Nil(t, err)
}

func TestUserRepository_CreateWithValidData(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	repository := NewUserRepository(db)

	if err != nil {
		t.Error(err)
	}

	user, err := entity.NewUser("any_name", "any_valid_email@gmail.com", "any_password")

	err = repository.Create(user)
	assert.Nil(t, err)
	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.Id).Error
	assert.Nil(t, err)

	assert.Equal(t, userFound.Id, user.Id)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Password, user.Password)
	assert.Equal(t, userFound.Email, user.Email)
}
