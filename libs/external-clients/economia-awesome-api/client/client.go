package client

import (
	"context"
	gorequest "libs/shared/go-request"
	"net/http"
	"time"

	outputDTO "libs/services/acl/dtos/economia-awesome-api/output"
)

// Client represents a client for the Economia Awesome API.
type Client struct {
	// ctx is the context for API requests.
	ctx context.Context
	// baseURL is the base URL for the API.
	baseURL string
	// httpClient is the client used to make HTTP requests.
	httpClient *http.Client
	// Timeout for the API requests.
	Timeout time.Duration
}

// NewClient creates and returns a new Client instance with default settings.
func NewClient() *Client {
	return &Client{
		ctx:        context.Background(),
		baseURL:    "https://economia.awesomeapi.com.br",
		httpClient: &http.Client{},
		Timeout:    900 * time.Millisecond,
	}
}

// GetExchangeRate fetches the exchange rate for the given exchangeRateName from the Economia Awesome API.
func (c *Client) GetExchangeRate(exchangeRateName string) (outputDTO.CurrencyInfoMapDTO, error) {
	pathParams := []string{"json", "last", exchangeRateName}
	headers := map[string]string{"Content-Type": "application/json"}
	req, err := gorequest.CreateRequest(
		c.ctx,
		c.baseURL,
		pathParams,
		nil,
		nil,
		headers,
		http.MethodGet,
	)
	if err != nil {
		return nil, err
	}

	var apiOutput outputDTO.CurrencyInfoMapDTO
	err = gorequest.SendRequest(c.ctx, req, c.httpClient, &apiOutput, c.Timeout)
	if err != nil {
		return outputDTO.CurrencyInfoMapDTO{}, err
	}
	return apiOutput, nil
}
