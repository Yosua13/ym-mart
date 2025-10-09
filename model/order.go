package model

import "time"

type Order struct {
	OrderID       int         `json:"order_id"`
	UserID        int         `json:"user_id"`
	InvoiceNumber string      `json:"invoice_number"`
	TotalAmount   float64     `json:"total_amount"`
	Status        string      `json:"status"`
	OrderItems    []OrderItem `json:"order_items"`
	CreatedAt     time.Time   `json:"created_at"`
}

type OrderItem struct {
	OrderItemID     int     `json:"order_item_id"`
	OrderID         int     `json:"order_id"`
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
	Quantity        int     `json:"quantity"`
}
