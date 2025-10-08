package model

import "time"

type Store struct {
	StoreID   int       `json:"store_id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ProductID   int       `json:"product_id"`
	StoreID     int       `json:"store_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}

type Cart struct {
	CartID    int        `json:"cart_id"`
	UserID    int        `json:"user_id"`
	CartItems []CartItem `json:"cart_items"`
	CreatedAt time.Time  `json:"created_at"`
}

type CartItem struct {
	CartItemID int     `json:"cart_item_id"`
	CartID     int     `json:"cart_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Product    Product `json:"product"`
}

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
