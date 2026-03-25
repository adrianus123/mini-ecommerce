package dto

type PaymentRequest struct {
	OrderID       string  `json:"order_id"`
	UserID        string  `json:"user_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type PaymentResponse struct {
	Success     bool   `json:"success"`
	PaymentCode string `json:"payment_code"`
	Message     string `json:"message"`
}
