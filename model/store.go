package model

import "time"

type Store struct {
	StoreID   int       `json:"store_id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}
