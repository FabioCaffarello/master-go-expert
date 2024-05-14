package gorequest

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RequestTestSuite struct {
	suite.Suite
}

func TestRequestSuite(t *testing.T) {
	suite.Run(t, new(RequestTestSuite))
}

func (suite *RequestTestSuite) TestParseBaseUrlWhenBaseUrlIsValid() {
	result, err := parseBaseURL("https://dummie.com")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "https", result.Scheme)
	assert.Equal(suite.T(), "dummie.com", result.Host)
}

func (suite *RequestTestSuite) TestParseBaseUrlWhenBaseUrlIsInvalid() {
	_, err := parseBaseURL("invalid-url")
	assert.NotNil(suite.T(), err)
}

func (suite *RequestTestSuite) TestBuildURLWhenThereIsPathParamsAndQueryParams() {
	expected := "https://dummie.com/param1/param2?query1=value1&query2=value2"
	result, err := buildURL("https://dummie.com", []string{"param1", "param2"}, map[string]string{"query1": "value1", "query2": "value2"})
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
}

func (suite *RequestTestSuite) TestBuildURLWhenThereIsPathParams() {
	expected := "https://dummie.com/param1/param2"
	result, err := buildURL("https://dummie.com/", []string{"param1", "param2"}, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
}

func (suite *RequestTestSuite) TestBuildURLWhenThereIsQueryParams() {
	expected := "https://dummie.com?query1=value1&query2=value2"
	result, err := buildURL("https://dummie.com", nil, map[string]string{"query1": "value1", "query2": "value2"})
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
}

func (suite *RequestTestSuite) TestBuildURLWhenThereIsNoPathParamsAndQueryParams() {
	expected := "https://dummie.com"
	result, err := buildURL("https://dummie.com", nil, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
}

func (suite *RequestTestSuite) TestBuildURLWhenBaseURLIsInvalid() {
	_, err := buildURL("invalid-url", nil, map[string]string{"query1": "value1", "query2": "value2"})
	assert.NotNil(suite.T(), err)
}

func (suite *RequestTestSuite) TestMarshalBodyWhenBodyIsNil() {
	contentType := "application/json"
	result, err := marshalBody(nil, contentType)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []byte{}, result)
}

func (suite *RequestTestSuite) TestMarshalBodyWhenBodyIsJsonContentType() {
	contentType := "application/json"
	result, err := marshalBody(map[string]string{"key": "value"}, contentType)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "{\"key\":\"value\"}", string(result))
}

func (suite *RequestTestSuite) TestMarshalBodyWhenBodyIsXmlContentType() {
	type XMLData struct {
		XMLName xml.Name `xml:"key"`
		Value   string   `xml:",chardata"`
	}

	contentType := "application/xml"
	xmlData := XMLData{Value: "value"}
	result, err := marshalBody(xmlData, contentType)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "<key>value</key>", string(result))
}

func (suite *RequestTestSuite) TestMarshalBodyWhenBodyIsFormUrlEncodedContentType() {
	contentType := "application/x-www-form-urlencoded"
	body := url.Values{"key": {"value"}}
	result, err := marshalBody(body, contentType)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "key=value", string(result))
}

func (suite *RequestTestSuite) TestMarshalBodyWhenBodyIsUnsupportedContentType() {
	contentType := "unsupported-content-type"
	_, err := marshalBody("body", contentType)
	assert.NotNil(suite.T(), err)
}

func (suite *RequestTestSuite) TestSetHeaders() {
	headers := map[string]string{"Content-Type": "application/json", "key": "value"}
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "https://dummie.com", nil)
	setHeaders(req, headers)
	assert.Equal(suite.T(), "application/json", req.Header.Get("Content-Type"))
	assert.Equal(suite.T(), "value", req.Header.Get("key"))
}

func (suite *RequestTestSuite) TestSetQueryParams() {
	parsedURL, _ := url.Parse("https://dummie.com")
	queryParams := map[string]string{"query1": "value1", "query2": "value2"}
	setQueryParams(parsedURL, queryParams)
	assert.Equal(suite.T(), "query1=value1&query2=value2", parsedURL.RawQuery)
}

func (suite *RequestTestSuite) TestSetPathParams() {
	parsedURL, _ := url.Parse("https://dummie.com/")
	pathParams := []string{"param1", "param2"}
	err := setPathParams(parsedURL, pathParams)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "https://dummie.com/param1/param2", parsedURL.String())
}

func (suite *RequestTestSuite) TestSetHeaderDefaultContentType() {
	headers := map[string]string{"key": "value"}
	contentType := "application/json"
	setHeaderDefaultContentType(headers, contentType)
	assert.Equal(suite.T(), "application/json", headers["Content-Type"])
}

func (suite *RequestTestSuite) TestGetContentType() {
	headers := map[string]string{"Content-Type": "application/json"}
	contentType := getContentType(headers)
	assert.Equal(suite.T(), "application/json", contentType)
}

func (suite *RequestTestSuite) TestGetContentTypeWhenContentTypeIsNotPresent() {
	headers := map[string]string{"key": "value"}
	contentType := getContentType(headers)
	assert.Equal(suite.T(), defaultContentType, contentType)
}

func (suite *RequestTestSuite) TestCreateRequestWhenBaseUrlIsInvalid() {
	method := "GET"
	headers := map[string]string{"Content-Type": "application/json"}
	url := "invalid-url"
	_, err := CreateRequest(context.Background(), url, nil, nil, nil, headers, method)
	assert.NotNil(suite.T(), err)
}

func (suite *RequestTestSuite) TestCreateRequestWhenMarshallingBodyFails() {
	method := "GET"
	headers := map[string]string{"Content-Type": "unsupported-content-type"}
	url := "https://dummie.com"
	_, err := CreateRequest(context.Background(), url, nil, nil, "body", headers, method)
	assert.NotNil(suite.T(), err)
}

func (suite *RequestTestSuite) TestCreateRequestWhenMethodIsGet() {
	method := "GET"
	url := "https://dummie.com"
	headers := map[string]string{"Content-Type": "application/json"}
	req, err := CreateRequest(context.Background(), url, nil, nil, nil, headers, method)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.NoBody, req.Body)
	assert.Equal(suite.T(), "GET", req.Method)
	assert.Equal(suite.T(), "https://dummie.com", req.URL.String())
	assert.Equal(suite.T(), "application/json", req.Header.Get("Content-Type"))
}

func (suite *RequestTestSuite) TestCreateRequestWhenMethodIsPost() {
	method := "POST"
	url := "https://dummie.com"
	headers := map[string]string{"Content-Type": "application/json"}
	body := map[string]string{"key": "value"}
	req, err := CreateRequest(context.Background(), url, nil, nil, body, headers, method)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "POST", req.Method)
	assert.Equal(suite.T(), "https://dummie.com", req.URL.String())
	assert.Equal(suite.T(), "application/json", req.Header.Get("Content-Type"))
}

type MockResponse struct {
	Message string `json:"message"`
}

func TestSendRequest_Success(t *testing.T) {
	expectedResult := MockResponse{Message: "success"}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResult)
	}))
	defer server.Close()

	ctx := context.Background()
	req, err := CreateRequest(ctx, server.URL, nil, nil, nil, map[string]string{"Content-Type": "application/json"}, http.MethodGet)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	var result MockResponse
	err = SendRequest(ctx, req, server.Client(), &result, time.Second)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if result.Message != expectedResult.Message {
		t.Errorf("Expected %s, got %s", expectedResult.Message, result.Message)
	}
}

func TestSendRequest_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // simulate a long processing time
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx := context.Background()
	req, err := CreateRequest(ctx, server.URL, nil, nil, nil, map[string]string{"Content-Type": "application/json"}, http.MethodGet)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	var result MockResponse
	err = SendRequest(ctx, req, server.Client(), &result, 1*time.Second) // timeout set to 1 second
	if err == nil || err.Error() != "HTTP request timed out" {
		t.Errorf("Expected timeout error, got %v", err)
	}
}

func TestSendRequest_HttpError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}))
	defer server.Close()

	ctx := context.Background()
	req, err := CreateRequest(ctx, server.URL, nil, nil, nil, map[string]string{"Content-Type": "application/json"}, http.MethodGet)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	var result MockResponse
	err = SendRequest(ctx, req, server.Client(), &result, time.Second)
	if err == nil || err.Error() != "HTTP request failed" {
		t.Errorf("Expected HTTP request failed error, got %v", err)
	}
}

func TestSendRequest_FailedToDecode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not a valid json"))
	}))
	defer server.Close()

	ctx := context.Background()
	req, err := CreateRequest(ctx, server.URL, nil, nil, nil, map[string]string{"Content-Type": "application/json"}, http.MethodGet)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	var result MockResponse
	err = SendRequest(ctx, req, server.Client(), &result, time.Second)
	if err == nil || err.Error() != "failed to decode response body" {
		t.Errorf("Expected failed to decode response body error, got %v", err)
	}
}
