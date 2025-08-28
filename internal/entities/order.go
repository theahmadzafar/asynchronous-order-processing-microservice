package entities

type Order struct {
	OrderId string   `json:"order_id"`
	UserId  string   `json:"user_id"`
	Items   []string `json:"items"`
	Total   int64    `json:"total"`
}
