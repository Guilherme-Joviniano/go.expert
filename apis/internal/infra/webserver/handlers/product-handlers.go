package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Guilherme-Joviniano/go.expert/apis/internal/dto"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	database "github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/database/schemas"
	public_entity "github.com/Guilherme-Joviniano/go.expert/apis/pkg/entity"
	"github.com/Guilherme-Joviniano/go.expert/apis/pkg/util"
	"github.com/go-chi/chi/v5"
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

func (h *ProductHandler) GetProduct(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductService.GetById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product

	productId, err := public_entity.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productBody, err := util.RequestToTypeAdapter[dto.UpdateProductInput](req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.Id = productId
	product.Name = productBody.Name
	product.Price = productBody.Price

	updatedProduct, err := h.ProductService.Update(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedProduct)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := public_entity.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: check different errors types to switch between http responses
	err = h.ProductService.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
