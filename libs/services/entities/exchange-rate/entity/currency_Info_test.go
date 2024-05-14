package exchangerateentity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CurrencyInfoEntityTestSuite struct {
	suite.Suite
}

func TestCurrencyInfoEntityTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyInfoEntityTestSuite))
}

func (suite *CurrencyInfoEntityTestSuite) TestNewExchangeRate() {
	code := "USD"
	codeIn := "BRL"
	name := "Dollar"
	high := 5.5
	low := 5.4
	varBid := 5.45
	pctChange := 0.01
	bid := 5.45
	ask := 5.46
	timestamp := int64(1626889200)
	createDate := "2021-07-21 00:00:00"

	currencyInfo, err := NewExchangeRate(
		code,
		codeIn,
		name,
		high,
		low,
		varBid,
		pctChange,
		bid,
		ask,
		timestamp,
		createDate,
	)

	suite.NoError(err)
	suite.NotNil(currencyInfo)
	suite.Equal(code, currencyInfo.Code)
	suite.Equal(codeIn, currencyInfo.CodeIn)
	suite.Equal(name, currencyInfo.Name)
	suite.Equal(high, currencyInfo.High)
	suite.Equal(low, currencyInfo.Low)
	suite.Equal(varBid, currencyInfo.VarBid)
	suite.Equal(pctChange, currencyInfo.PctChange)
	suite.Equal(bid, currencyInfo.Bid)
	suite.Equal(ask, currencyInfo.Ask)
	suite.Equal(timestamp, currencyInfo.Timestamp)
	suite.Equal(createDate, currencyInfo.CreateDate.Format("2006-01-02 15:04:05"))
	suite.Equal("55c1e4be-5777-5cbb-9db9-39218834db87", currencyInfo.ID)
}

func (suite *CurrencyInfoEntityTestSuite) TestNewExchangeRateInvalidWhenCodeIsEmpty() {
	code := ""
	codeIn := "BRL"
	name := "Dollar"
	high := 5.5
	low := 5.4
	varBid := 5.45
	pctChange := 0.01
	bid := 5.45
	ask := 5.46
	timestamp := int64(1626889200)
	createDate := "2021-07-21 00:00:00"

	currencyInfo, err := NewExchangeRate(
		code,
		codeIn,
		name,
		high,
		low,
		varBid,
		pctChange,
		bid,
		ask,
		timestamp,
		createDate,
	)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), currencyInfo)
}

func (suite *CurrencyInfoEntityTestSuite) TestNewExchangeRateInvalidWhenCodeInIsEmpty() {
	code := "USD"
	codeIn := ""
	name := "Dollar"
	high := 5.5
	low := 5.4
	varBid := 5.45
	pctChange := 0.01
	bid := 5.45
	ask := 5.46
	timestamp := int64(1626889200)
	createDate := "2021-07-21 00:00:00"

	currencyInfo, err := NewExchangeRate(
		code,
		codeIn,
		name,
		high,
		low,
		varBid,
		pctChange,
		bid,
		ask,
		timestamp,
		createDate,
	)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), currencyInfo)
}

func (suite *CurrencyInfoEntityTestSuite) TestNewExchangeRateInvalidWhenBidIsZero() {
	code := "USD"
	codeIn := "BRL"
	name := "Dollar"
	high := 5.5
	low := 5.4
	varBid := 5.45
	pctChange := 0.01
	bid := 0.0
	ask := 5.46
	timestamp := int64(1626889200)
	createDate := "2021-07-21 00:00:00"

	currencyInfo, err := NewExchangeRate(
		code,
		codeIn,
		name,
		high,
		low,
		varBid,
		pctChange,
		bid,
		ask,
		timestamp,
		createDate,
	)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), currencyInfo)
}

func (suite *CurrencyInfoEntityTestSuite) TestNewExchangeRateInvalidWhenTimestamp() {
	code := "USD"
	codeIn := "BRL"
	name := "Dollar"
	high := 5.5
	low := 5.4
	varBid := 5.45
	pctChange := 0.01
	bid := 5.45
	ask := 5.46
	timestamp := int64(0)
	createDate := "2021-07-21 00:00:00"

	currencyInfo, err := NewExchangeRate(
		code,
		codeIn,
		name,
		high,
		low,
		varBid,
		pctChange,
		bid,
		ask,
		timestamp,
		createDate,
	)

	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), currencyInfo)
}
