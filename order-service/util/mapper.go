package util

import (
	"fmt"
	"order-service/internal/delivery/http/dto"
	"order-service/internal/entity"
	"time"
)

func ConstructOrderEntity(totalPrice float64, userID, paymentMethod, status string) entity.Order {
	now := time.Now()
	orderCode := "TRX-" + fmt.Sprint(now.UnixMilli())

	return entity.Order{
		OrderCode:     orderCode,
		UserID:        userID,
		TotalPrice:    totalPrice,
		PaymentMethod: paymentMethod,
		Status:        status,
	}
}

func ConstructOrderItemEntity(orderId, qty uint, product dto.Product) entity.OrderItem {
	subtotal := product.Price * float64(qty)

	return entity.OrderItem{
		OrderID:      uint(orderId),
		ProductID:    product.ID,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Qty:          qty,
		Subtotal:     subtotal,
	}
}

func ConstructPaymentRequest(order entity.Order) dto.PaymentRequest {
	return dto.PaymentRequest{
		OrderID:       order.OrderCode,
		UserID:        order.UserID,
		Amount:        order.TotalPrice,
		PaymentMethod: order.PaymentMethod,
	}
}

func ConstructEventItem(req dto.Item) entity.Item {
	return entity.Item{
		ProductID: uint(req.ProductID),
		Qty:       req.Qty,
	}
}

func ConstructEvent(owner, action string, value []entity.Item) entity.Event {
	return entity.Event{
		Owner:  owner,
		Action: action,
		Value:  value,
	}
}
