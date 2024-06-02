package gosd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceDiscoveryTestSuite struct {
	suite.Suite
}

func TestServiceDiscoveryTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceDiscoveryTestSuite))
}

func (suite *ServiceDiscoveryTestSuite) SetupTest() {
	os.Clearenv()
}

func (suite *ServiceDiscoveryTestSuite) TestNewServiceDiscovery() {
	sd := NewServiceDiscovery(true)
	assert.NotNil(suite.T(), sd)
	assert.True(suite.T(), sd.UsesHostAddr)
	assert.NotNil(suite.T(), sd.services)
}

func (suite *ServiceDiscoveryTestSuite) TestRegisterService() {
	sd := NewServiceDiscovery(true)
	sd.RegisterService("exchange-rate-api", "8080")
	assert.Equal(suite.T(), "http://exchange-rate-api:8080", sd.services["exchange-rate-api"])
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURL() {
	sd := NewServiceDiscovery(true)
	sd.RegisterService("exchange-rate-api", "8080")
	url, err := sd.GetBaseURL("exchange-rate-api")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "http://exchange-rate-api:8080", url)
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURLNoHostAddr() {
	sd := NewServiceDiscovery(false)
	sd.RegisterService("exchange-rate-api", "8080")
	url, err := sd.GetBaseURL("exchange-rate-api")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "http://localhost:8080", url)
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURLFromEnv() {
	sd := NewServiceDiscovery(true)
	os.Setenv("EXCHANGE_RATE_API_HOST", "exchange-rate-api")
	os.Setenv("EXCHANGE_RATE_API_PORT", "8080")
	url, err := sd.GetBaseURLFromEnv("EXCHANGE_RATE_API")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "http://exchange-rate-api:8080", url)
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURLFromEnvError() {
	sd := NewServiceDiscovery(true)
	os.Setenv("EXCHANGE_RATE_API_HOST", "")
	os.Setenv("EXCHANGE_RATE_API_PORT", "")
	url, err := sd.GetBaseURLFromEnv("EXCHANGE_RATE_API")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "", url)
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURLFromEnvNoHostAddr() {
	sd := NewServiceDiscovery(false)
	os.Setenv("EXCHANGE_RATE_API_HOST", "exchange-rate-api")
	os.Setenv("EXCHANGE_RATE_API_PORT", "8080")
	url, err := sd.GetBaseURLFromEnv("EXCHANGE_RATE_API")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "http://localhost:8080", url)
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURLFromEnvNoHostAddrError() {
	sd := NewServiceDiscovery(false)
	os.Setenv("EXCHANGE_RATE_API_HOST", "")
	os.Setenv("EXCHANGE_RATE_API_PORT", "")
	url, err := sd.GetBaseURLFromEnv("EXCHANGE_RATE_API")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "", url)
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURLNotFound() {
	sd := NewServiceDiscovery(true)
	_, err := sd.GetBaseURL("exchange-rate-api")
	assert.Error(suite.T(), err)
}

func (suite *ServiceDiscoveryTestSuite) TestGetBaseURLFromEnvNotFound() {
	sd := NewServiceDiscovery(true)
	_, err := sd.GetBaseURLFromEnv("EXCHANGE_RATE_API")
	assert.Error(suite.T(), err)
}



