package exchangerateentity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	dateStringLayout = "2006-01-02 15:04:05"
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
	high := "5.5"
	low := "5.4"
	varBid := "5.45"
	pctChange := "0.01"
	bid := "5.45"
	ask := "5.46"
	timestamp := "1626889200"
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

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), currencyInfo)
	assert.Equal(suite.T(), code, currencyInfo.Code)
	assert.Equal(suite.T(), codeIn, currencyInfo.CodeIn)
	assert.Equal(suite.T(), name, currencyInfo.Name)
	assert.Equal(suite.T(), 5.5, currencyInfo.High)
	assert.Equal(suite.T(), 5.4, currencyInfo.Low)
	assert.Equal(suite.T(), 5.45, currencyInfo.VarBid)
	assert.Equal(suite.T(), 0.01, currencyInfo.PctChange)
	assert.Equal(suite.T(), 5.45, currencyInfo.Bid)
	assert.Equal(suite.T(), 5.46, currencyInfo.Ask)
	assert.Equal(suite.T(), int64(1626889200), currencyInfo.Timestamp)
	suite.Equal(createDate, currencyInfo.CreateDate.Format("2006-01-02 15:04:05"))
	suite.Equal("83da6030-ab1f-5b6c-8b07-7bac10f85dbc", currencyInfo.ID)
}

func (suite *CurrencyInfoEntityTestSuite) TestNewExchangeRateInvalidWhenCodeIsEmpty() {
	code := ""
	codeIn := "BRL"
	name := "Dollar"
	high := "5.5"
	low := "5.4"
	varBid := "5.45"
	pctChange := "0.01"
	bid := "5.45"
	ask := "5.46"
	timestamp := "1626889200"
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
	high := "5.5"
	low := "5.4"
	varBid := "5.45"
	pctChange := "0.01"
	bid := "5.45"
	ask := "5.46"
	timestamp := "1626889200"
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
	high := "5.5"
	low := "5.4"
	varBid := "5.45"
	pctChange := "0.01"
	bid := "0"
	ask := "5.46"
	timestamp := "1626889200"
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
	high := "5.5"
	low := "5.4"
	varBid := "5.45"
	pctChange := "0.01"
	bid := "5.45"
	ask := "5.46"
	timestamp := "invalid"
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

func (suite *CurrencyInfoEntityTestSuite) TestMapToCurrencyInfoEntity() {
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
	createDate, _ := time.Parse(dateStringLayout, "2021-07-21 00:00:00")

	document := map[string]interface{}{
		"_id":        "83da6030-ab1f-5b6c-8b07-7bac10f85dbc",
		"code":       code,
		"codein":     codeIn,
		"name":       name,
		"high":       high,
		"low":        low,
		"varBid":     varBid,
		"pctChange":  pctChange,
		"bid":        bid,
		"ask":        ask,
		"timestamp":  timestamp,
		"create_date": createDate,
	}

	var result CurrencyInfo
	err := MapToCurrencyInfoEntity(document, &result)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), code, result.Code)
	assert.Equal(suite.T(), codeIn, result.CodeIn)
	assert.Equal(suite.T(), name, result.Name)
	assert.Equal(suite.T(), high, result.High)
	assert.Equal(suite.T(), low, result.Low)
	assert.Equal(suite.T(), varBid, result.VarBid)
	assert.Equal(suite.T(), pctChange, result.PctChange)
	assert.Equal(suite.T(), bid, result.Bid)
	assert.Equal(suite.T(), ask, result.Ask)
	assert.Equal(suite.T(), timestamp, result.Timestamp)
	assert.Equal(suite.T(), createDate, result.CreateDate)
}
