package client

import "context"

// import (
// 	"libs/shared/go-request"
// 	"libs/services/acl/dtos/economia-awesome-api"
// )

type Client struct {
	ctx     context.Context
	baseURL string
}

func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "https://economia-awesome-api.com",
	}
}


func (c *Client) GetExchangeRate() (string, error) {
	
	return "1.23", nil
}