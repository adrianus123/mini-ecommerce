package handler

import (
	"order-service/internal/delivery/http/dto"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(usecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (h *OrderHandler) Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello Order Service"})
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest

	c.BindJSON(&req)

	resp, err := h.usecase.CreateOrder(req, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": resp})
}
