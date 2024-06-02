package sqliterepository

import (
	"libs/resources/database/in-memory/sqlite-client/client"
	entity "libs/services/entities/exchange-rate/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SQLiteDBExchangeRateRepositoryTestSuite struct {
	suite.Suite
	client     *client.Client
	repository *ExchangeRateRepository
}

func TestSQLiteDBExchangeRateRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SQLiteDBExchangeRateRepositoryTestSuite))
}

func (suite *SQLiteDBExchangeRateRepositoryTestSuite) SetupTest() {
	databasePath := ":memory:"
	cl, err := client.NewClient(databasePath)
	assert.NoError(suite.T(), err)

	suite.client = cl
	suite.repository = NewExchangeRateRepository(databasePath, cl)
}

func (suite *SQLiteDBExchangeRateRepositoryTestSuite) TearDownTest() {
	suite.client.Close()
}

func (suite *SQLiteDBExchangeRateRepositoryTestSuite) TestCreateTable() {
	err := suite.repository.createTable()
	assert.NoError(suite.T(), err)

	query := "SELECT name FROM sqlite_master WHERE type='table' AND name='exchange_rates'"
	row := suite.client.QueryRow(query)

	var tableName string
	err = row.Scan(&tableName)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "exchange_rates", tableName)
}
func (suite *SQLiteDBExchangeRateRepositoryTestSuite) TestSave() {
	err := suite.repository.createTable()
	assert.NoError(suite.T(), err)

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

	currencyInfo, _ := entity.NewExchangeRate(
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

	err = suite.repository.Save(currencyInfo)
	assert.NoError(suite.T(), err)

	query := "SELECT id, code, codeIn, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date FROM exchange_rates WHERE id = ?"
	row := suite.client.QueryRow(query, currencyInfo.ID)

	var retrievedCurrencyInfo entity.CurrencyInfo
	err = row.Scan(
		&retrievedCurrencyInfo.ID,
		&retrievedCurrencyInfo.Code,
		&retrievedCurrencyInfo.CodeIn,
		&retrievedCurrencyInfo.Name,
		&retrievedCurrencyInfo.High,
		&retrievedCurrencyInfo.Low,
		&retrievedCurrencyInfo.VarBid,
		&retrievedCurrencyInfo.PctChange,
		&retrievedCurrencyInfo.Bid,
		&retrievedCurrencyInfo.Ask,
		&retrievedCurrencyInfo.Timestamp,
		&retrievedCurrencyInfo.CreateDate,
	)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), currencyInfo, &retrievedCurrencyInfo)
}
