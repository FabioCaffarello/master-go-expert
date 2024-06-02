# Exchange Rate Repository

The `exchange-rate` library provides a repository to handle CRUD operations for exchange rate entities using an in-memory document database.

## Overview

This package includes the following main components:
- `ExchangeRateRepository`: A struct that provides methods to perform CRUD operations on exchange rate entities within the in-memory document database.

## Features

The main functionalities provided by the package include:
- Ensuring the collection exists.
- Saving exchange rate entities to the collection.
- Finding exchange rate entities by various criteria.
- Deleting exchange rate entities from the collection.

## Types

- **ExchangeRateRepository**: Provides methods to interact with exchange rate entities in the in-memory document database.

## Functions

### ExchangeRateRepository Functions

- `NewExchangeRateRepository(database string, client *client.Client) *ExchangeRateRepository`: Creates and returns a new `ExchangeRateRepository` instance.
- `Save(currencyInfo *entity.CurrencyInfo) error`: Saves the given currency info entity into the collection.
- `FindAll() ([]*entity.CurrencyInfo, error)`: Retrieves all exchange rate entities from the collection.
- `FindByID(id string) (*entity.CurrencyInfo, error)`: Retrieves a single exchange rate entity by its ID from the collection.
- `Find(code string, codeIn string) ([]*entity.CurrencyInfo, error)`: Retrieves exchange rate entities by their code and codeIn from the collection.
- `Delete(id string) error`: Removes a single exchange rate entity by its ID from the collection.

## Usage

### Creating a New Repository

```go
import (
	"libs/resources/database/in-memory/go-doc-db-client/client"
	entity "libs/services/entities/exchange-rate/entity"
	repository "libs/services/infrastructure/database/repositories/exchange-rate/in-memory/go-doc-db/repository"
)

dbClient := client.NewClient(database.NewInMemoryDocBD("myDatabase"))
repository := repository.NewExchangeRateRepository("myDatabase", dbClient)
```

### Saving a Currency Info

```go
currencyInfo := &entity.CurrencyInfo{
    ID:     "12345",
    Code:   "USD",
    CodeIn: "BRL",
    Rate:   5.42,
}

err := repository.Save(currencyInfo)
if err != nil {
    log.Fatal(err)
}
```

### Finding All Currency Infos

```go
currencyInfos, err := repository.FindAll()
if err != nil {
    log.Fatal(err)
}
fmt.Println(currencyInfos)
```

### Finding a Currency Info by ID

```go
currencyInfo, err := repository.FindByID("12345")
if err != nil {
    log.Fatal(err)
}
fmt.Println(currencyInfo)
```

### Finding Currency Infos by Code

```go
currencyInfos, err := repository.Find("USD", "BRL")
if err != nil {
    log.Fatal(err)
}
fmt.Println(currencyInfos)
```

### Deleting a Currency Info by ID

```go
err := repository.Delete("12345")
if err != nil {
    log.Fatal(err)
}
```
