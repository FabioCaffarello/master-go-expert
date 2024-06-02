package client

import (
	outputDTO "libs/services/acl/dtos/economia-awesome-api/output"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EconomiaAwesomeAPIClientTestSuite struct {
	suite.Suite
}

func TestEconomiaAwesomeAPIClientSuite(t *testing.T) {
	suite.Run(t, new(EconomiaAwesomeAPIClientTestSuite))
}

func (suite *EconomiaAwesomeAPIClientTestSuite) TestNewClient() {
	client := NewClient()
	assert.NotNil(suite.T(), client)
}

func (suite *EconomiaAwesomeAPIClientTestSuite) TestGetExchangeRate() {
	client := NewClient()
	result, err := client.GetExchangeRate("EUR-BRL")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.NotEmpty(suite.T(), result)
	assert.IsType(suite.T(), outputDTO.CurrencyInfoMapDTO{}, result)
}
