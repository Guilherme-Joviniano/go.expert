package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("any_name", float32(19.99))

	assert.NotNil(t, product)
	assert.Equal(t, product.Name, "any_name")
	assert.Equal(t, product.Price, float32(19.99))
	assert.Nil(t, err)
}

func TestNewProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", float32(19.99))

	assert.Nil(t, product)
	assert.Error(t, err, ErrNameIsRequired)
}

func TestNewProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("any_name", float32(0.0))

	assert.Nil(t, product)
	assert.Error(t, err, ErrPriceIsRequired)
}
