package repository

import (
	"order-service/internal/entity"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepository {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) CreateOrder(order *entity.Order) error {
	result := r.db.Create(order)
	return result.Error
}

func (r *OrderRepo) CreateOrderItem(orderItem *entity.OrderItem) error {
	result := r.db.Create(orderItem)
	return result.Error
}

func (r *OrderRepo) SaveOrder(order *entity.Order) error {
	result := r.db.Save(order)
	return result.Error
}
