package database

import (
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) FindByEmail(emailIdentifier string) (*entity.User, error) {
	var user entity.User

	if err := repo.DB.Where("email = ?", emailIdentifier).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Create(user *entity.User) error {
	return repo.DB.Create(user).Error
}
