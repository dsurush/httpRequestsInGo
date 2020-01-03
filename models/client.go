package models

// Client "Данные которые хранятся клиенте"
type Client struct {
	ID           string `json:"id,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	SurName      string `json:"surName,omitempty"`
	Address      string `json:"address,omitempty"`
	Email        string `json:"email,omitempty"`
	PhoneNumbers string `json:"phoneNumbers,omitempty"`
}

// CreateNewClient "Структура добавлении нового клиента"
type CreateNewClient struct {
	FirstName    string `json:"firstName"`
	SurName      string `json:"surName"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	PhoneNumbers string `json:"phoneNumbers"`
}

// AddClientResponse "напишем какой структуры будет запрос на добавление нового клиента"
type AddClientResponse struct {
	Error       bool    `json:"error,omitempty"`
	Description string  `json:"description,omitempty"`
	Client      *Client `json:"client"`
}

// AddClientRequeste "напишем какой структуры будет ответ на добавление нового клиента"
type AddClientRequeste struct {
	Client CreateNewClient `json:"client"`
}
