package database

import (
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emailIdentifier string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	GetById(id string) (*entity.Product, error)
	List(page, limit int, sort string) ([]entity.Product, error)
	Delete(id string) error
	Update(product *entity.Product) (*entity.Product, error)
}
