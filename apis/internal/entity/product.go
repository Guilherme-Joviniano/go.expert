package entity

import (
	"errors"
	"time"

	"github.com/Guilherme-Joviniano/go.expert/apis/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
	ErrInvalidId       = errors.New("invalid id")
)

type Product struct {
	Id        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt string    `json:"created_at"`
}

func NewProduct(name string, price float32) (*Product, error) {
	product := &Product{
		Id:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().String(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.Id.String() == "" {
		return ErrIdIsRequired
	}
	if _, err := entity.ParseID(p.Id.String()); err != nil {
		return ErrInvalidId
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price < float32(0.0) {
		return ErrInvalidPrice
	}
	if p.Price == 0.0 {
		return ErrPriceIsRequired
	}
	return nil
}
