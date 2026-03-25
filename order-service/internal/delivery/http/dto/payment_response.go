package dto

type PaymentResponse struct {
	Data Payment `json:"data"`
}

type Payment struct {
	Success     bool   `json:"success"`
	PaymentCode string `json:"payment_code"`
	Message     string `json:"message"`
}
