package repository

import (
	"payment-service/internal/entity"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentRepository {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) CreatePayment(payment *entity.Payment) error {
	result := r.db.Create(payment)
	return result.Error
}

func (r *PaymentRepo) Save(payment *entity.Payment) error {
	result := r.db.Save(payment)
	return result.Error
}
