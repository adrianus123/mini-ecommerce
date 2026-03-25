package port

import "order-service/internal/delivery/http/dto"

type PaymentService interface {
	CreatePayment(req dto.PaymentRequest) (*dto.Payment, error)
}
