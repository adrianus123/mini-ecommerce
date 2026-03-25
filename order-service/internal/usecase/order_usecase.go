package usecase

import (
	"errors"
	"fmt"
	"log"
	"order-service/internal/application/port"
	"order-service/internal/delivery/http/dto"
	"order-service/internal/entity"
	"order-service/internal/repository"
	"order-service/util"

	"github.com/gin-gonic/gin"
)

type OrderUsecase struct {
	repo           repository.OrderRepository
	productService port.ProductService
	paymentService port.PaymentService
	kafkaUsecase   *PublishEventUsecase
}

func NewOrderUsecase(orderRepo repository.OrderRepository, productService port.ProductService, paymentService port.PaymentService, kafkaUsecase *PublishEventUsecase) *OrderUsecase {
	return &OrderUsecase{
		repo:           orderRepo,
		productService: productService,
		paymentService: paymentService,
		kafkaUsecase:   kafkaUsecase,
	}
}

func (u *OrderUsecase) CreateOrder(req dto.CreateOrderRequest, c *gin.Context) (string, error) {
	totalPrice := 0.0
	products := map[int]dto.Product{}
	eventItem := []entity.Item{}

	for _, item := range req.Orders {
		product, err := u.productService.GetProduct(item.ProductID)

		if err != nil {
			return "", err
		}

		if product.Qty == 0 {
			return "", errors.New("This item is empty")
		}

		if product.Qty < item.Qty {
			return "", errors.New("Insufficient number of items")
		}

		totalPrice += product.Price * float64(item.Qty)
		products[item.ProductID] = *product
		eventItem = append(eventItem, util.ConstructEventItem(item))
	}

	userID := util.GetUser(c)
	order := util.ConstructOrderEntity(totalPrice, userID, "TRANSFER", "PENDING")
	err := u.repo.CreateOrder(&order)
	if err != nil {
		return "", err
	}

	for _, item := range req.Orders {
		orderItem := util.ConstructOrderItemEntity(order.ID, item.Qty, products[item.ProductID])
		err := u.repo.CreateOrderItem(&orderItem)
		if err != nil {
			return "", err
		}
	}

	paymentReq := util.ConstructPaymentRequest(order)
	payment, err := u.paymentService.CreatePayment(paymentReq)
	if err != nil {
		return "", fmt.Errorf("Error payment process: %w", err)
	}

	log.Println("Payment: ", payment)

	var message string
	if payment.Success {
		order.Status = "PAID"
		message = "Order Successful"

		err := u.kafkaUsecase.Execute(util.ConstructEvent("order-service", "UPDATE", eventItem))
		if err != nil {
			log.Println("Failed publish event to kafka: ", err)
		}
	} else {
		order.Status = "FAILED"
		message = "Order Failed"
	}

	err = u.repo.SaveOrder(&order)
	if err != nil {
		return "", fmt.Errorf("Error while update status order: %w", err)
	}

	return message, nil
}
