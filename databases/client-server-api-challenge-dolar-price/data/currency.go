package data

import (
	"context"
	"errors"
	pkg_http "github.com/Guilherme-Joviniano/go-currency-api/pkg/http"
	"github.com/Guilherme-Joviniano/go-currency-api/types"
	"gorm.io/gorm"
	"time"
)

type Currency struct {
	ID             int `gorm:"primaryKey"`
	Value          string
	CurrencyCode   CurrencyCode
	CurrencyCodeId int
	gorm.Model
}

type CurrencyCode struct {
	ID   int `gorm:"primaryKey"`
	Code string
	gorm.Model
}

type CurrencyServiceInterface interface {
	Save(ctx context.Context, currency *Currency) error
	Get(ctx context.Context, code string) (*types.HttpResponse[types.DolarAPI], error)
}

type CurrencyAdapter struct {
	db *gorm.DB
}

type CurrencyHttpAdapter struct{}

func (c *CurrencyAdapter) Save(ctx context.Context, currency *Currency) error {

	err := c.db.Model(&Currency{}).WithContext(ctx).Create(&currency).Error

	select {
	case <-ctx.Done():
		return errors.New("timeout db exceed time")
	default:
		return err
	}
}

func (c *CurrencyAdapter) Get(ctx context.Context, code string) (*types.HttpResponse[*types.DolarAPI], error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	result, err := pkg_http.FetchHttpThirdPartyCurrencyAPI(ctx)

	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("Exceeded timeout reached")
	default:
		return result, nil
	}
}

func NewCurrencyAdapter(database *gorm.DB) *CurrencyAdapter {
	return &CurrencyAdapter{
		db: database,
	}
}
