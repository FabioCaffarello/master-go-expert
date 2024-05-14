package gorequest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var (
	DefaultHTTPClient = &http.Client{
		Timeout: 500 * time.Millisecond,
	}
	contentType = "application/json"
)

func CreateRequest(
	ctx context.Context,
	method string,
	urlStr string,
	body interface{},
	headers map[string]string,
	queryParams map[string]string,
) (*http.Request, error) {
	parsedURL, err := buildURL(urlStr, queryParams)
	if err != nil {
		return nil, err
	}

	requestBody, err := marshalBody(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, parsedURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	setHeaders(req, headers)

	return req, nil
}

func buildURL(urlStr string, queryParams map[string]string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}
	if queryParams != nil {
		query := parsedURL.Query()
		for key, value := range queryParams {
			query.Add(key, value)
		}
		parsedURL.RawQuery = query.Encode()
	}
	return parsedURL.String(), nil
}

func marshalBody(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	return requestBody, nil
}

func setHeaders(req *http.Request, headers map[string]string) {
	req.Header.Set("Content-Type", contentType)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
