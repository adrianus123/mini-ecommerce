package usecase

import (
	"fmt"
	"math/rand"
	"payment-service/internal/delivery/http/dto"
	"payment-service/internal/repository"
	"payment-service/util"
)

type PaymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{repo: paymentRepo}
}

func (u *PaymentUsecase) CreatePayment(req dto.PaymentRequest) (dto.PaymentResponse, error) {
	payment := util.ConstructPayment(req)

	err := u.repo.CreatePayment(&payment)
	if err != nil {
		return dto.PaymentResponse{}, fmt.Errorf("Error save payment to database: %w", err)
	}

	// random status payment
	success := rand.Intn(100) < 80

	var resp dto.PaymentResponse
	if success {
		resp = util.ConstructPaymentResponse(success, payment.PaymentCode, "Payment Successful")
		payment.Status = "SUCCESS"
	} else {
		resp = util.ConstructPaymentResponse(success, payment.PaymentCode, "Payment Failed")
		payment.Status = "FAILED"
	}

	err = u.repo.Save(&payment)
	if err != nil {
		return dto.PaymentResponse{}, fmt.Errorf("Error while update status payment: %w", err)
	}

	return resp, nil
}
