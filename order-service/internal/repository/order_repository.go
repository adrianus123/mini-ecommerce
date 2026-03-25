package repository

import "order-service/internal/entity"

type OrderRepository interface {
	CreateOrder(order *entity.Order) error
	CreateOrderItem(orderItem *entity.OrderItem) error
	SaveOrder(order *entity.Order) error
}
