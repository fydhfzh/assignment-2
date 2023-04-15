package dto

import "github.com/fydhfzh/assignment-2/entity"

type NewOrderRequest struct {
	OrderedAt    string           `json:"orderedAt"`
	CustomerName string           `json:"customerName"`
	Items        []NewItemRequest `json:"items"`
}

type GetOrderRequest struct {
	OrderedAt    string           `json:"orderedAt"`
	CustomerName string           `json:"customerName"`
	Items        []NewItemRequest `json:"items"`
}

type UpdateOrderRequest struct {
	OrderedAt    string           `json:"orderedAt"`
	CustomerName string           `json:"customerName"`
	Items        []NewItemRequest `json:"items"`
}

type OrderResponse struct {
	OrderId      int    `json:"id,omitempty"`
	CreatedAt	string 	`json:"created_at,omitempty"`
	UpdatedAt	string 	`json:"updated_at,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	Items        []entity.Item `json:"items,omitempty"`
}

type NewOrderResponse struct {
	StatusCode int `json:"code,omitempty"`
	Data OrderResponse `json:"data,omitempty"`
}

type GetOrderResponse struct{
	Data OrderResponse `json:"data,omitempty"`
}

type UpdateOrderResponse struct{
	StatusCode int `json:"code,omitempty"`
	Data OrderResponse `json:"data,omitempty"`
}

