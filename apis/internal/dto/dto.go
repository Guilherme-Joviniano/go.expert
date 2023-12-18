package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
type UpdateProductInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

