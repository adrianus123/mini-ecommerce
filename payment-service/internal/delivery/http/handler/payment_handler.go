package handler

import (
	"net/http"
	"payment-service/internal/delivery/http/dto"
	"payment-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	usecase *usecase.PaymentUsecase
}

func NewPaymentHandler(u *usecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{usecase: u}
}

func (h *PaymentHandler) Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello Payment Service"})
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req dto.PaymentRequest

	c.BindJSON(&req)

	resp, err := h.usecase.CreatePayment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
