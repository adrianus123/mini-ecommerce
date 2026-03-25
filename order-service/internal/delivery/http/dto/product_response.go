package dto

import (
	"time"
)

type ProductResponse struct {
	Data Product `json:"data"`
}

type Product struct {
	ID        uint      `json:"ID"`
	Name      string    `json:"Name"`
	Price     float64   `json:"Price"`
	Qty       uint      `json:"Qty"`
	CreatedAt time.Time `json:"CreatedAt"`
}
