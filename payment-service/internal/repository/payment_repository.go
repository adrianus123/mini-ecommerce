package repository

import "payment-service/internal/entity"

type PaymentRepository interface {
	CreatePayment(p *entity.Payment) error
	Save(p *entity.Payment) error
}
