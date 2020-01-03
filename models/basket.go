package models

//Basket "-Robins-Pobins"
type Basket struct {
	ID      string  `json:"id,omitempty"`
	Product Product `json:"product,omitempty"`
}
