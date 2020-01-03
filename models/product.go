package models

// Product "создаем саму структуру Продукта"
type Product struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Cost    int    `json:"cost"`
	Picture string `json:"picture,omitempty"`
}

// AddProductRequest "Бла бла бла"
type AddProductRequest struct {
	Products Product `json:"animals,omitempty"`
}

// AddProductResponse "Будет отвечать за буль буль"
type AddProductResponse struct {
	Error       bool     `json:"error,omitempty"`
	Description string   `json:"description,omitempty"`
	Products    *Product `json:"products,omitempty"`
}
