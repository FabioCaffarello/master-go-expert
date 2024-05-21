package gorequest

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	defaultContentType = "application/json"
)

// parseBaseURL parses the given base URL and returns a parsed *url.URL or an error if the URL is invalid.
func parseBaseURL(baseURL string) (*url.URL, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid base URL: missing scheme or host")
	}
	return parsedURL, nil
}

// buildURL constructs a full URL with the given base URL, path parameters, and query parameters.
// Returns the full URL as a string or an error if any component is invalid.
func buildURL(baseURL string, pathParams []string, queryParams map[string]string) (string, error) {
	// Parse the base URL
	parsedURL, err := parseBaseURL(baseURL)
	if err != nil {
		return "", err
	}

	// Add path parameters
	if pathParams != nil {
		err = setPathParams(parsedURL, pathParams)
		if err != nil {
			return "", err
		}
	}

	// Add query parameters
	if queryParams != nil {
		setQueryParams(parsedURL, queryParams)
	}

	return parsedURL.String(), nil
}

// marshalBody marshals the given body into bytes based on the specified content type.
// Supports JSON, XML, and URL-encoded forms. Returns the marshaled bytes or an error if the content type is unsupported.
func marshalBody(body interface{}, contentType string) ([]byte, error) {
	if body == nil {
		return []byte{}, nil
	}

	switch contentType {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	case "application/x-www-form-urlencoded":
		return []byte(body.(url.Values).Encode()), nil
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
}

// setHeaders sets the provided headers on the given HTTP request.
func setHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// setQueryParams adds the given query parameters to the URL.
func setQueryParams(parsedURL *url.URL, queryParams map[string]string) {
	query := parsedURL.Query()
	for key, value := range queryParams {
		query.Set(key, value)
	}
	parsedURL.RawQuery = query.Encode()
}

// setPathParams appends the given path parameters to the URL path.
func setPathParams(parsedURL *url.URL, pathParams []string) error {
	var err error
	joinedPathParams := strings.Join(pathParams, "/")
	parsedURL.Path, err = url.JoinPath(parsedURL.Path, joinedPathParams)
	if err != nil {
		return errors.New("failed to join path parameters")
	}
	return nil
}

// getContentType retrieves the Content-Type from the headers or sets and returns the default content type.
func getContentType(headers map[string]string) string {
	if contentType, ok := headers["Content-Type"]; ok {
		return contentType
	}
	setHeaderDefaultContentType(headers, defaultContentType)
	return defaultContentType
}

// setHeaderDefaultContentType sets the Content-Type header to the default content type.
func setHeaderDefaultContentType(headers map[string]string, contentType string) {
	headers["Content-Type"] = contentType
}

// CreateRequest creates an HTTP request with the given parameters.
// It builds the URL, marshals the body, and sets the headers. Returns the constructed *http.Request or an error.
func CreateRequest(
	ctx context.Context,
	baseUrl string,
	pathParams []string,
	queryParams map[string]string,
	body interface{},
	headers map[string]string,
	method string,
) (*http.Request, error) {
	parsedURL, err := buildURL(baseUrl, pathParams, queryParams)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	contentType := getContentType(headers)

	requestBody, err := marshalBody(body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, parsedURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	setHeaders(req, headers)

	return req, nil
}

// SendRequest sends the given HTTP request using the provided client.
// It waits for the response or times out after the specified duration. The response body is decoded into the result parameter.
// Returns an error if the request fails, times out, or the response status is not 2xx.
func SendRequest(
	ctx context.Context,
	req *http.Request,
	client *http.Client,
	result interface{},
	timeout time.Duration,
) error {
	type responseResult struct {
		resp *http.Response
		err  error
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	resultCh := make(chan responseResult, 1)

	go func() {
		resp, err := client.Do(req)
		resultCh <- responseResult{resp: resp, err: err}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("HTTP request timed out: %w", ctx.Err())
	case res := <-resultCh:
		if res.err != nil {
			return fmt.Errorf("failed to send HTTP request: %w", res.err)
		}
		defer res.resp.Body.Close()

		if res.resp.StatusCode < http.StatusOK || res.resp.StatusCode >= http.StatusMultipleChoices {
			return fmt.Errorf("HTTP request failed: %s", res.resp.Status)
		}

		if err := json.NewDecoder(res.resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response body: %w", err)
		}
		return nil
	}
}
