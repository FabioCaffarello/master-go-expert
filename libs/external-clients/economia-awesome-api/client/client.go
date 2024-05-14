package client

import (
	"context"
	gorequest "libs/shared/go-request"
	"time"

	// inputDTO "libs/services/acl/dtos/economia-awesome-api/input"
	outputDTO "libs/services/acl/dtos/exchange-rate/output"
)

type Client struct {
	// ctx is the context for API requests.
	ctx context.Context
	// baseURL is the base URL for the API.
	baseURL string
}

func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "https://economia-awesome-api.com",
	}
}

func (c *Client) GetExchangeRate(exchangeRateName string) (outputDTO.ExchangeRatesDTO, error) {
	// json/last/EUR-BRL
	pathParams := []string{"json", "last", exchangeRateName}
	req, err := gorequest.CreateRequest(
		c.ctx,
		c.baseURL,
		pathParams,
		nil,
		nil,
		nil,
		"GET",
	)
	if err != nil {
		return nil, err
	}
	var apiOutput outputDTO.ExchangeRatesDTO
	err = gorequest.SendRequest(c.ctx, req, nil, &apiOutput, 500*time.Millisecond)
	if err != nil {
		return outputDTO.ExchangeRatesDTO{}, err
	}
	return apiOutput, nil
}
