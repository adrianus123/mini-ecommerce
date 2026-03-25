package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order-service/internal/delivery/http/dto"
)

type ProductClient struct {
	BaseURL string
}

func NewProductClient(baseUrl string) *ProductClient {
	return &ProductClient{
		BaseURL: baseUrl,
	}
}

func (c *ProductClient) GetProduct(productID int) (*dto.Product, error) {
	url := fmt.Sprintf("%s/products/%d", c.BaseURL, productID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to call product service: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Product service returned status %d: %s", resp.StatusCode, string(body))
	}

	var productResp dto.ProductResponse
	if err := json.Unmarshal(body, &productResp); err != nil {
		return nil, fmt.Errorf("Failed to decode product response: %w", err)
	}

	return &productResp.Data, nil
}
