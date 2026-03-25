package dto

type CreateOrderRequest struct {
	Orders []Item `json:"orders"`
}

type Item struct {
	ProductID int  `json:"product_id"`
	Qty       uint `json:"qty"`
}
