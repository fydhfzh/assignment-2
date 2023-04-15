package entity

type Item struct {
	ItemId      int    `json:"id"`
	ItemCode    string `json:"itemcode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"orderid"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}