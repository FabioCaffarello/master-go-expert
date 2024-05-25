package main

import (
	"context"
	outputDTO "libs/services/acl/dtos/exchange-rate/output"
	gorequest "libs/shared/go-request"
	gosd "libs/shared/go-sd/service_discovery"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var (
	useHostAddress = false
	serviceName    = "api-gateway"
	port           = "5000"
)

// 

func main() {
	sd := gosd.NewServiceDiscovery(useHostAddress)
	gatewayClient := NewClient(sd)
	currencyInfo, err := gatewayClient.GetCotacoes(context.Background())
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("currencyInfo: %v", currencyInfo)
	tmpl := template.New("cotacoes")
	for _, currency := range currencyInfo {
		tmpl, _ = tmpl.Parse("Dólar: {{.Bid}}")
		tmpl.Execute(os.Stdout, currency)
	}
}

type Client struct {
	sd         gosd.ServiceDiscoveryInterface
	httpClient *http.Client
	Timeout    time.Duration
}

func NewClient(
	sd gosd.ServiceDiscoveryInterface,
) *Client {
	sd.RegisterService(serviceName, port)
	return &Client{
		sd:         sd,
		httpClient: &http.Client{},
		Timeout:    900 * time.Millisecond,
	}
}

func (c *Client) getBaseURL() (string, error) {
	return c.sd.GetBaseURL(serviceName)
}

func (c *Client) GetCotacoes(ctx context.Context) (outputDTO.ExchangeRatesDTO, error) {
	url, err := c.getBaseURL()
	if err != nil {
		return nil, err
	}
	pathParams := []string{"cotacoes"}
	headers := map[string]string{"Content-Type": "application/json"}
	req, err := gorequest.CreateRequest(
		ctx,
		url,
		pathParams,
		nil,
		nil,
		headers,
		http.MethodGet,
	)
	if err != nil {
		return nil, err
	}

	var apiOutput outputDTO.ExchangeRatesDTO
	err = gorequest.SendRequest(ctx, req, c.httpClient, &apiOutput, c.Timeout)
	if err != nil {
		return outputDTO.ExchangeRatesDTO{}, err
	}
	log.Printf("apiOutput: %v", apiOutput)
	return apiOutput, nil
}
