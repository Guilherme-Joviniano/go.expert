package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewProductRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	repository := NewProductRepository(db)
	assert.NotNil(t, repository)
}

func TestProductRepository_GetByIdWithNonRegisteredIdValue(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	product, err := repository.GetById("invalid_id")

	assert.Nil(t, product)
	assert.Contains(t, err.Error(), "record not found")
}

func TestProductRepository_GetByIdWithAValidId(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("any_product", float32(19.99))

	db.Create(&product)

	productFound, err := repository.GetById(product.Id.String())

	assert.Equal(t, productFound.Id.String(), product.Id.String())
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
	assert.Nil(t, err)
}
func TestProductRepository_FindByDescWithNoPaginationOrLimit(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	products := make([]entity.Product, 30)

	for i := 1; i <= 30; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("product_%d", i), rand.Float32()*100)

		if err != nil {
			t.Error(err)
		}

		db.Create(&product)

		products[i-1] = *product
	}

	productsFound, err := repository.List(0, 0, "desc")

	assert.Len(t, productsFound, 30)

	for idx := range products {
		assert.Equal(t, products[idx].Id.String(), productsFound[idx].Id.String())
		assert.Equal(t, products[idx].Name, productsFound[idx].Name)
		assert.Equal(t, products[idx].Price, productsFound[idx].Price)
	}

	assert.Nil(t, err)
}
func TestProductRepository_FindByAscWithNoPaginationOrLimit(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	products := make([]entity.Product, 30)

	for i := 1; i <= 30; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("product_%d", i), rand.Float32()*100)

		if err != nil {
			t.Error(err)
		}

		db.Create(&product)

		products[i-1] = *product
	}

	productsFound, err := repository.List(0, 0, "asc")

	assert.Len(t, productsFound, 30)

	for idx := range products {
		assert.Equal(t, products[idx].Id.String(), productsFound[idx].Id.String())
		assert.Equal(t, products[idx].Name, productsFound[idx].Name)
		assert.Equal(t, products[idx].Price, productsFound[idx].Price)
	}

	assert.Nil(t, err)
}
func TestProductRepository_FindWithPaginationAndLimit(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	products := make([]entity.Product, 10)

	for i := 1; i <= 10; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("product_%d", i), rand.Float32()*100)

		if err != nil {
			t.Error(err)
		}

		db.Create(&product)

		products[i-1] = *product
	}

	productsFound, err := repository.List(1, 10, "asc")

	assert.Len(t, productsFound, 10)

	for idx := range products {
		assert.Equal(t, products[idx].Id.String(), productsFound[idx].Id.String())
		assert.Equal(t, products[idx].Name, productsFound[idx].Name)
		assert.Equal(t, products[idx].Price, productsFound[idx].Price)
	}

	assert.Nil(t, err)
}
func TestProductRepository_UpdateWithInvalidData(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct(fmt.Sprintf("product_%d", 1), rand.Float32()*100)

	if err != nil {
		t.Error(err)
	}

	db.Create(product)

	product.Price = float32(-22.05)

	updatedProduct, err := repository.Update(product)

	assert.Nil(t, updatedProduct)
	assert.Contains(t, err.Error(), "invalid price")
}
func TestProductRepository_UpdateWithValidData(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct(fmt.Sprintf("product_%d", 1), rand.Float32()*100)

	if err != nil {
		t.Error(err)
	}

	db.Create(product)

	product.Price = float32(22.05)

	updatedProduct, err := repository.Update(product)

	assert.Equal(t, updatedProduct.Id.String(), product.Id.String())
	assert.Equal(t, updatedProduct.Name, product.Name)
	assert.Equal(t, updatedProduct.Price, product.Price)
	assert.Nil(t, err)
}

func TestProductRepository_CreateWithValidData(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	repository := NewProductRepository(db)

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct(fmt.Sprintf("product_%d", 1), rand.Float32()*100)

	if err != nil {
		t.Error(err)
	}

	err = repository.Create(product)

	assert.Nil(t, err)
}
