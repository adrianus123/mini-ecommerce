package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order-service/internal/delivery/http/dto"
)

type PaymentClient struct {
	BaseURL string
}

func NewPaymentClient(baseUrl string) *PaymentClient {
	return &PaymentClient{
		BaseURL: baseUrl,
	}
}

func (c *PaymentClient) CreatePayment(req dto.PaymentRequest) (*dto.Payment, error) {
	url := fmt.Sprintf("%s/payments/create", c.BaseURL)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling request to json: :%w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Failed to call product service: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("payment service returned error status: %d, body: %s", resp.StatusCode, string(body))
	}

	var paymentResp dto.PaymentResponse
	err = json.Unmarshal(body, &paymentResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(body))
	}

	return &paymentResp.Data, nil
}
