package handlers

import (
	"net/http"

	"github.com/Guilherme-Joviniano/go.expert/apis/internal/dto"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	database "github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/database/schemas"
	"github.com/Guilherme-Joviniano/go.expert/apis/pkg/util"
)

type ProductHandler struct {
	ProductService database.ProductInterface
}

func NewProductHandler(service database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	product, err := util.RequestToTypeAdapter[dto.CreateProductInput](req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// should go to uc layer

	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductService.Create(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
