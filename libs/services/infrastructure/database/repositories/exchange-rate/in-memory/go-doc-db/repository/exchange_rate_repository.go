package godocdbrepository

import (
	"libs/resources/database/in-memory/go-doc-db-client/client"
	entity "libs/services/entities/exchange-rate/entity"
	"log"
)

var (
	collectionName = "currency-info"
)

type ExchangeRateRepository struct {
	database          string
	client            *client.Client
	collectionName    string
	collectionCreated bool
}

func NewExchangeRateRepository(
	database string,
	client *client.Client,
) (*ExchangeRateRepository, error) {
	exchangeRateRepository := &ExchangeRateRepository{
		database:          database,
		client:            client,
		collectionName:    collectionName,
		collectionCreated: false,
	}
	return exchangeRateRepository, nil
}

// init initializes the repository by ensuring the collection exists
func (r *ExchangeRateRepository) init() {
	if err := r.createCollectionIfNotExists(); err != nil {
		log.Fatal(err)
	}
}

// CreateCollectionIfNotExists checks if the collection is already created, if not then creates it
func (r *ExchangeRateRepository) createCollectionIfNotExists() error {
	if !r.collectionCreated {
		err := r.client.CreateCollection(r.collectionName)
		if err != nil {
			log.Printf("Error creating collection: %v", err)
			return err
		}
		r.collectionCreated = true
		return nil
	}
	return nil
}

func (r *ExchangeRateRepository) Save(currencyInfo *entity.CurrencyInfo) error {
	log.Printf("Saving exchange rate to collection: %v", r.collectionName)
	r.init()
	currencyInfoMap := currencyInfo.ToMap()
	err := r.client.InsertOne(r.collectionName, currencyInfoMap)
	if err != nil {
		log.Printf("Error saving exchange rate: %v", err)
		return err
	}
	return nil
}

func (r *ExchangeRateRepository) FindAll() ([]*entity.CurrencyInfo, error) {
	log.Printf("Finding all exchange rates from collection: %v", r.collectionName)
	r.init()
	documents, err := r.client.FindAll(r.collectionName)
	if err != nil {
		log.Printf("Error finding all exchange rates: %v", err)
		return nil, err
	}
	currencyInfos := make([]*entity.CurrencyInfo, len(documents))
	for i, document := range documents {
		result, err := entity.MapToCurrencyInfoEntity(document)
		if err != nil {
			log.Printf("Error mapping document to entity: %v", err)
			return nil, err
		}
		currencyInfos[i] = result
	}
	return currencyInfos, nil
}

func (r *ExchangeRateRepository) FindByID(id string) (*entity.CurrencyInfo, error) {
	log.Printf("Finding exchange rate by ID from collection: %v", r.collectionName)
	r.init()
	document, err := r.client.FindOne(r.collectionName, id)
	if err != nil {
		log.Printf("Error finding exchange rate by ID: %v", err)
		return nil, err
	}
	result, err := entity.MapToCurrencyInfoEntity(document)
	if err != nil {
		log.Printf("Error mapping document to entity: %v", err)
		return nil, err
	}
	return result, nil
}

func (r *ExchangeRateRepository) Find(code string, codeIn string) ([]*entity.CurrencyInfo, error) {
	log.Printf("Finding exchange rate by code from collection: %v", r.collectionName)
	r.init()
	queryFilter := map[string]interface{}{
		"code":   code,
		"codeIn": codeIn,
	}

	documents, err := r.client.Find(r.collectionName, queryFilter)
	if err != nil {
		log.Printf("Error finding exchange rate by code: %v", err)
		return nil, err
	}

	currencyInfos := make([]*entity.CurrencyInfo, len(documents))
	for i, document := range documents {
		result, err := entity.MapToCurrencyInfoEntity(document)
		if err != nil {
			log.Printf("Error mapping document to entity: %v", err)
			return nil, err
		}
		currencyInfos[i] = result
	}
	return currencyInfos, nil
}
