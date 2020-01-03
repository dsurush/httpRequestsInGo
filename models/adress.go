package models

//Adress "создание структуры"
type Adress struct {
	ID     string `json:"id,omitempty"`
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
	Home   string `json:"home,omitempty"`
}
