package client

import (
	"context"
	outputDTO "libs/services/acl/dtos/exchange-rate/output"
	gorequest "libs/shared/go-request"
	gosd "libs/shared/go-sd/service_discovery"
	"log"
	"net/http"
	"time"
)

var (
	serviceName = "exchange-rate-api"
	port        = "8080"
)

type Client struct {
	ctx        context.Context
	sd         gosd.ServiceDiscoveryInterface
	httpClient *http.Client
	Timeout    time.Duration
}

func NewClient(
	sd gosd.ServiceDiscoveryInterface,
) *Client {
	sd.RegisterService(serviceName, port)
	return &Client{
		ctx:        context.Background(),
		sd:         sd,
		httpClient: &http.Client{},
		Timeout:    900 * time.Millisecond,
	}
}

func (c *Client) getBaseURL() (string, error) {
	return c.sd.GetBaseURL(serviceName)
}

func (c *Client) GetExchangeRate(code string, codeIn string) (outputDTO.ExchangeRatesDTO, error) {
	url, err := c.getBaseURL()
	log.Printf("url: %v", url)
	if err != nil {
		return nil, err
	}
	queryParams := map[string]string{
		"code":    code,
		"code_in": codeIn,
	}
	pathParams := []string{"cotacoes"}
	headers := map[string]string{"Content-Type": "application/json"}
	req, err := gorequest.CreateRequest(
		c.ctx,
		url,
		pathParams,
		queryParams,
		nil,
		headers,
		http.MethodGet,
	)
	if err != nil {
		return nil, err
	}

	var apiOutput outputDTO.ExchangeRatesDTO
	err = gorequest.SendRequest(c.ctx, req, c.httpClient, &apiOutput, c.Timeout)
	if err != nil {
		return outputDTO.ExchangeRatesDTO{}, err
	}
	log.Printf("apiOutput: %v", apiOutput)
	return apiOutput, nil
}
