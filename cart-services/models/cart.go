package models

type Cart struct {
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	Id    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
