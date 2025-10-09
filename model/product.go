package model

import "time"

type Product struct {
	ProductID   int       `json:"product_id"`
	StoreID     int       `json:"store_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}
