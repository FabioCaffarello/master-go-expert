package godocdbrepository

import (
	"testing"

	"libs/resources/database/in-memory/go-doc-db-client/client"
	"libs/resources/database/in-memory/go-doc-db/database"
	entity "libs/services/entities/exchange-rate/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoDocDBExchangeRateRepositoryTestSuite struct {
	suite.Suite
	client           *client.Client
	databaseName     string
	collectionName   string
	currencyInfoData *entity.CurrencyInfo
}

func TestGoDocDBExchangeRateRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GoDocDBExchangeRateRepositoryTestSuite))
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) SetupTest() {
	var err error
	suite.databaseName = "test-database"
	suite.collectionName = "currency-info"
	db := database.NewInMemoryDocBD(suite.databaseName)
	suite.client = client.NewClient(db)
	suite.currencyInfoData, err = entity.NewExchangeRate(
		"USD",
		"BRL",
		"Dollar",
		"5.5",
		"5.4",
		"5.45",
		"0.01",
		"5.45",
		"5.46",
		"1626889200",
		"2021-07-21 00:00:00",
	)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), suite.currencyInfoData)
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TearDownTest() {
	suite.client = nil
	suite.databaseName = ""
	suite.collectionName = ""
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestNewGoDocDBExchangeRateRepository() {
	testCase := struct {
		database        string
		client          *client.Client
		collectionName  string
	}{
		database:        suite.databaseName,
		client:          suite.client,
		collectionName:  suite.collectionName,
	}

	repository, err := NewExchangeRateRepository(
		testCase.database,
		testCase.client,
	)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), repository)
	assert.Equal(suite.T(), testCase.database, repository.database)
	assert.Equal(suite.T(), testCase.client, repository.client)
	assert.Equal(suite.T(), testCase.collectionName, repository.collectionName)
	assert.False(suite.T(), repository.collectionCreated)
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestInit() {
	repository, _ := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)

	repository.init()
	assert.Equal(suite.T(), true, repository.collectionCreated)
	assert.Equal(suite.T(), []string{suite.collectionName}, suite.client.ListCollections())
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestCreateCollectionIfNotExists() {
	repository, _ := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)

	err := repository.createCollectionIfNotExists()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), true, repository.collectionCreated)
	assert.Equal(suite.T(), []string{suite.collectionName}, suite.client.ListCollections())
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestCreateCollectionIfNotExistsWhenCollectionAlreadyExists() {
	repository, _ := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)

	err := repository.createCollectionIfNotExists()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), true, repository.collectionCreated)
	assert.Equal(suite.T(), []string{suite.collectionName}, suite.client.ListCollections())

	err = repository.createCollectionIfNotExists()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), true, repository.collectionCreated)
	assert.Equal(suite.T(), []string{suite.collectionName}, suite.client.ListCollections())
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestSave() {
	repository, err := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)
	assert.Nil(suite.T(), err)

	repository.init()
	assert.Equal(suite.T(), true, repository.collectionCreated)

	err = repository.Save(suite.currencyInfoData)
	assert.Nil(suite.T(), err)

	results, err := suite.client.FindAll(suite.collectionName)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestSaveWhenCollectionNotCreated() {
	repository, err := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)
	assert.Nil(suite.T(), err)

	err = repository.Save(suite.currencyInfoData)
	assert.Nil(suite.T(), err)

	results, err := suite.client.FindAll(suite.collectionName)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestSaveWhenAlreadyexists() {
	repository, err := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)
	assert.Nil(suite.T(), err)

	repository.init()
	assert.Equal(suite.T(), true, repository.collectionCreated)

	err = repository.Save(suite.currencyInfoData)
	assert.Nil(suite.T(), err)

	results, err := suite.client.FindAll(suite.collectionName)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))

	err = repository.Save(suite.currencyInfoData)
	assert.NotNil(suite.T(), err)

	results, err = suite.client.FindAll(suite.collectionName)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestFindAll() {
	repository, err := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)
	assert.Nil(suite.T(), err)

	repository.init()
	assert.Equal(suite.T(), true, repository.collectionCreated)

	err = repository.Save(suite.currencyInfoData)
	assert.Nil(suite.T(), err)

	results, err := repository.FindAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))

	assert.Equal(suite.T(), suite.currencyInfoData.ID, results[0].ID)
	assert.Equal(suite.T(), suite.currencyInfoData.Code, results[0].Code)
	assert.Equal(suite.T(), suite.currencyInfoData.CodeIn, results[0].CodeIn)
	assert.Equal(suite.T(), suite.currencyInfoData.Name, results[0].Name)
	assert.Equal(suite.T(), suite.currencyInfoData.High, results[0].High)
	assert.Equal(suite.T(), suite.currencyInfoData.Low, results[0].Low)
	assert.Equal(suite.T(), suite.currencyInfoData.VarBid, results[0].VarBid)
	assert.Equal(suite.T(), suite.currencyInfoData.PctChange, results[0].PctChange)
	assert.Equal(suite.T(), suite.currencyInfoData.Bid, results[0].Bid)
	assert.Equal(suite.T(), suite.currencyInfoData.Ask, results[0].Ask)
	assert.Equal(suite.T(), suite.currencyInfoData.Timestamp, results[0].Timestamp)
	assert.Equal(suite.T(), suite.currencyInfoData.CreateDate, results[0].CreateDate)
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestFindByID() {
	repository, err := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)
	assert.Nil(suite.T(), err)

	repository.init()
	assert.Equal(suite.T(), true, repository.collectionCreated)

	err = repository.Save(suite.currencyInfoData)
	assert.Nil(suite.T(), err)

	result, err := repository.FindByID(suite.currencyInfoData.ID)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), suite.currencyInfoData.ID, result.ID)
	assert.Equal(suite.T(), suite.currencyInfoData.Code, result.Code)
	assert.Equal(suite.T(), suite.currencyInfoData.CodeIn, result.CodeIn)
	assert.Equal(suite.T(), suite.currencyInfoData.Name, result.Name)
	assert.Equal(suite.T(), suite.currencyInfoData.High, result.High)
	assert.Equal(suite.T(), suite.currencyInfoData.Low, result.Low)
	assert.Equal(suite.T(), suite.currencyInfoData.VarBid, result.VarBid)
	assert.Equal(suite.T(), suite.currencyInfoData.PctChange, result.PctChange)
	assert.Equal(suite.T(), suite.currencyInfoData.Bid, result.Bid)
	assert.Equal(suite.T(), suite.currencyInfoData.Ask, result.Ask)
	assert.Equal(suite.T(), suite.currencyInfoData.Timestamp, result.Timestamp)
	assert.Equal(suite.T(), suite.currencyInfoData.CreateDate, result.CreateDate)
}

func (suite *GoDocDBExchangeRateRepositoryTestSuite) TestFind() {
	repository, err := NewExchangeRateRepository(
		suite.databaseName,
		suite.client,
	)
	assert.Nil(suite.T(), err)

	repository.init()
	assert.Equal(suite.T(), true, repository.collectionCreated)

	err = repository.Save(suite.currencyInfoData)
	assert.Nil(suite.T(), err)

	results, err := repository.Find(suite.currencyInfoData.Code, suite.currencyInfoData.CodeIn)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(results))

	// assert.Equal(suite.T(), suite.currencyInfoData.ID, results[0].ID)
	// assert.Equal(suite.T(), suite.currencyInfoData.Code, results[0].Code)
	// assert.Equal(suite.T(), suite.currencyInfoData.CodeIn, results[0].CodeIn)
	// assert.Equal(suite.T(), suite.currencyInfoData.Name, results[0].Name)
	// assert.Equal(suite.T(), suite.currencyInfoData.High, results[0].High)
	// assert.Equal(suite.T(), suite.currencyInfoData.Low, results[0].Low)
	// assert.Equal(suite.T(), suite.currencyInfoData.VarBid, results[0].VarBid)
	// assert.Equal(suite.T(), suite.currencyInfoData.PctChange, results[0].PctChange)
	// assert.Equal(suite.T(), suite.currencyInfoData.Bid, results[0].Bid)
	// assert.Equal(suite.T(), suite.currencyInfoData.Ask, results[0].Ask)
	// assert.Equal(suite.T(), suite.currencyInfoData.Timestamp, results[0].Timestamp)
	// assert.Equal(suite.T(), suite.currencyInfoData.CreateDate, results[0].CreateDate)
}
