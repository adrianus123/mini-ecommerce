package util

import (
	"fmt"
	"payment-service/internal/delivery/http/dto"
	"payment-service/internal/entity"
	"time"
)

func ConstructPayment(req dto.PaymentRequest) entity.Payment {
	now := time.Now()
	paymentCode := fmt.Sprintf("PAY-%d", now.UnixMilli())

	return entity.Payment{
		PaymentCode:     paymentCode,
		OrderID:         req.OrderID,
		UserID:          req.UserID,
		Amount:          req.Amount,
		PaymentMethod:   req.PaymentMethod,
		PaymentProvider: GetPaymentProvider(req.PaymentMethod),
		Status:          "PENDING",
	}
}

func ConstructPaymentResponse(isSuccess bool, paymentCode, message string) dto.PaymentResponse {
	return dto.PaymentResponse{
		Success:     isSuccess,
		PaymentCode: paymentCode,
		Message:     message,
	}
}

func GetPaymentProvider(paymentMethod string) string {
	switch paymentMethod {
	case "TRANSFER":
		return "VIRTUAL ACCOUNT"
	case "DIGITAL":
		return "E-WALLET"
	default:
		return "PAYMENT GATEWAY"
	}
}
