package entity

type Order struct {
	OrderId      int    `json:"id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type OrderItem struct {
	OrderData Order  `json:"order"`
	Items     []Item `json:"items"`
}
