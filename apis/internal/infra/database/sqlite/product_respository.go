package database

import (
	"errors"

	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (repo *ProductRepository) GetById(id string) (*entity.Product, error) {
	var product entity.Product

	if err := repo.DB.Find(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if product.Name == "" {
		return nil, errors.New("record not found")
	}

	return &product, nil
}

func (repo *ProductRepository) Create(product *entity.Product) error {
	return repo.DB.Create(product).Error
}

func (repo *ProductRepository) Delete(id string) error {
	product, err := repo.GetById(id)

	if err != nil {
		return err
	}

	return repo.DB.Delete(product).Error
}

func (repo *ProductRepository) Update(product *entity.Product) (*entity.Product, error) {
	err := product.Validate()

	if err != nil {
		return nil, err
	}

	_, err = repo.GetById(product.Id.String())

	if err != nil {
		return nil, err
	}

	err = repo.DB.Save(product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductRepository) List(page, limit int, sort string) ([]entity.Product, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		return nil, errors.New("invalid sort type")
	}

	var products []entity.Product

	if page != 0 && limit != 0 {
		err := repo.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error

		if err != nil {
			return nil, err
		}

		return products, nil
	}

	err := repo.DB.Find(&products).Order("created_at " + sort).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}
