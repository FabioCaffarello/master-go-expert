package client

import (
	"context"
	"encoding/json"
	"fmt"
	outputDTO "libs/services/acl/dtos/exchange-rate/output"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExchangeRateAPIClientTestSuite struct {
	suite.Suite
	sd     *MockServiceDiscovery
	client *Client
	server *httptest.Server
}

func TestExchangeRateAPIClientTestSuite(t *testing.T) {
	suite.Run(t, new(ExchangeRateAPIClientTestSuite))
}

// Mocking the ServiceDiscovery
type MockServiceDiscovery struct {
	services map[string]string
}

func NewMockServiceDiscovery() *MockServiceDiscovery {
	return &MockServiceDiscovery{
		services: make(map[string]string),
	}
}

func (msd *MockServiceDiscovery) RegisterService(name, port string) {
	msd.services[name] = "http://localhost:" + port
}

func (msd *MockServiceDiscovery) GetBaseURL(serviceName string) (string, error) {
	if url, exists := msd.services[serviceName]; exists {
		return url, nil
	}
	return "", fmt.Errorf("service %s not found", serviceName)
}

func (msd *MockServiceDiscovery) GetBaseURLFromEnv(serviceName string) (string, error) {
	return msd.GetBaseURL(serviceName)
}

// SetupTest runs before each test
func (suite *ExchangeRateAPIClientTestSuite) SetupTest() {
	// Mock server to simulate the exchange rate API
	mockData1 := map[string]interface{}{
		"USDBRL": map[string]interface{}{
			"code":        "USD",
			"codein":      "BRL",
			"name":        "Dólar Americano/Real Brasileiro",
			"high":        float64(5.3004),
			"low":         float64(5.3004),
			"varBid":      float64(0.0001),
			"pctChange":   float64(0.0019),
			"bid":         float64(5.3004),
			"ask":         float64(5.3005),
			"timestamp":   int64(1619913600000),
			"create_date": time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	body, _ := json.Marshal(mockData1)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})
	suite.server = httptest.NewServer(handler)

	suite.sd = NewMockServiceDiscovery()
	suite.client = &Client{
		ctx:        context.Background(),
		sd:         suite.sd,
		httpClient: suite.server.Client(),
		Timeout:    900 * time.Millisecond,
	}

	// Override the service URL for testing
	suite.sd.RegisterService(serviceName, suite.server.URL[len("http://localhost:"):])
}

// TearDownTest runs after each test
func (suite *ExchangeRateAPIClientTestSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *ExchangeRateAPIClientTestSuite) TestNewClient() {
	client := NewClient(suite.sd)
	assert.NotNil(suite.T(), client)
	assert.NotNil(suite.T(), client.sd)
}

func (suite *ExchangeRateAPIClientTestSuite) TestGetBaseURL() {
	url, err := suite.client.getBaseURL()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "http://localhost:"+suite.server.URL[len("http://localhost:"):], url)
}

func (suite *ExchangeRateAPIClientTestSuite) TestGetExchangeRate() {
	result, err := suite.client.GetExchangeRate("USD", "BRL")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.NotEmpty(suite.T(), result)
	assert.IsType(suite.T(), outputDTO.ExchangeRatesDTO{}, result)
	assert.Len(suite.T(), result, 1)
	assert.Contains(suite.T(), result, "USDBRL")
	assert.Equal(suite.T(), "USD", result["USDBRL"].Code)
	assert.Equal(suite.T(), "BRL", result["USDBRL"].CodeIn)
}
