package entity

import (
	"errors"

	"github.com/Guilherme-Joviniano/go.expert/apis/pkg/entity"
)

var (
	ErrIdIsRequired = errors.New("id is required")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice = errors.New("invalid price")
	ErrInvalidId = errors.New("invalid id")
)

type Product struct {
	Id        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt string    `json:"created_at"`
}



