package client

import (
	"context"
	gorequest "libs/shared/go-request"
	"net/http"
	"time"

	outputDTO "libs/services/acl/dtos/exchange-rate/output"
)

type Client struct {
	// ctx is the context for API requests.
	ctx context.Context
	// baseURL is the base URL for the API.
	baseURL string
	// httpClient is the client used to make HTTP requests.
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		ctx:        context.Background(),
		baseURL:    "https://economia.awesomeapi.com.br",
		httpClient: &http.Client{},
	}
}

func (c *Client) GetExchangeRate(exchangeRateName string) (outputDTO.ExchangeRatesDTO, error) {
	pathParams := []string{"json", "last", exchangeRateName}
	headers := map[string]string{"Content-Type": "application/json"}
	req, err := gorequest.CreateRequest(
		c.ctx,
		c.baseURL,
		pathParams,
		nil,
		nil,
		headers,
		"GET",
	)
	if err != nil {
		return nil, err
	}
	var apiOutput outputDTO.ExchangeRatesDTO
	err = gorequest.SendRequest(c.ctx, req, c.httpClient, &apiOutput, 600*time.Millisecond)
	if err != nil {
		return outputDTO.ExchangeRatesDTO{}, err
	}
	return apiOutput, nil
}
