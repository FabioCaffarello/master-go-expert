package sqliterepository

import (
	"libs/resources/database/in-memory/sqlite-client/client"
	entity "libs/services/entities/exchange-rate/entity"
	"log"
)

var (
	schemaName = "currencyInfo"
)

type ExchangeRateRepository struct {
	database string
	client   *client.Client
	// collectionName    string
	// collectionCreated bool
}

func NewExchangeRateRepository(
	database string,
	client *client.Client,
) *ExchangeRateRepository {
	return &ExchangeRateRepository{
		database: database,
		client:   client,
		// collectionName:    collectionName,
		// collectionCreated: false,
	}
}

// CreateTable creates the exchange_rates table if it doesn't exist.
func (r *ExchangeRateRepository) createTable() error {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS exchange_rates (
        id TEXT PRIMARY KEY,
        code TEXT,
        codeIn TEXT,
        name TEXT,
        high REAL,
        low REAL,
        varBid REAL,
        pctChange REAL,
        bid REAL,
        ask REAL,
        timestamp INTEGER,
        create_date DATETIME
    )`
	return r.client.Exec(createTableQuery)
}

func (r *ExchangeRateRepository) Save(currencyInfo *entity.CurrencyInfo) error {
	if err := r.createTable(); err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}

	insertQuery := `
    INSERT INTO exchange_rates (id, code, codeIn, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    ON CONFLICT(id) DO UPDATE SET
        code = excluded.code,
        codeIn = excluded.codeIn,
        name = excluded.name,
        high = excluded.high,
        low = excluded.low,
        varBid = excluded.varBid,
        pctChange = excluded.pctChange,
        bid = excluded.bid,
        ask = excluded.ask,
        timestamp = excluded.timestamp,
        create_date = excluded.create_date
    `
	err := r.client.Exec(insertQuery, currencyInfo.GetEntityID(), currencyInfo.Code, currencyInfo.CodeIn, currencyInfo.Name, currencyInfo.High, currencyInfo.Low, currencyInfo.VarBid, currencyInfo.PctChange, currencyInfo.Bid, currencyInfo.Ask, currencyInfo.Timestamp, currencyInfo.CreateDate)
	return err
}
