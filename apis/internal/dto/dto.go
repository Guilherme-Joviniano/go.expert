package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
type UpdateProductInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name`
	Email    string `json:"email`
	Password string `json:"password`
}

type AuthenticationInput struct {
	Email    string `json:"email`
	Password string `json:"password`
}
