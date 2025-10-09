package model

import "time"

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
