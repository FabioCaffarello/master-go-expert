# exchange-rateentity Library

## Overview

The `exchange-rate` library provides a Go package to manage and manipulate exchange rate information for currency pairs entity. This library includes functionality to create, validate, and convert exchange rate data between different formats, and it also defines an interface for repositories to interact with this data.

## Usage

### CurrencyInfo

The `CurrencyInfo` struct represents the exchange rate entity information for a currency pair. 

#### Fields

- `ID`: Unique identifier for the exchange rate entity.
- `Code`: The base currency code.
- `CodeIn`: The target currency code.
- `Name`: The name of the currency pair.
- `High`: The highest rate observed.
- `Low`: The lowest rate observed.
- `VarBid`: The variation of the bid price.
- `PctChange`: The percentage change.
- `Bid`: The bid price.
- `Ask`: The ask price.
- `Timestamp`: The timestamp of the rate information.
- `CreateDate`: The date when the record was created.

### Creating a New CurrencyInfo Entity

You can create a new `CurrencyInfo` entity using the `NewExchangeRate` function, which validates the input and converts string values to appropriate data types.

```go
package main

import (
    "fmt"
    "log"
    "time"

    entity "libs/services/entities/exchange-rate"
)

func main() {
    exchangeRate, err := entity.NewExchangeRate(
        "USD",
		"BRL",
		"US Dollar/Brazilian Real",
        "5.4",
		"5.2",
		"0.2",
		"3.7",
		"5.3",
		"5.4",
        fmt.Sprintf("%d", time.Now().Unix()), time.Now().Format("2006-01-02 15:04:05"),
	)
    
    if err != nil {
        log.Fatalf("Error creating exchange rate: %v", err)
    }

    fmt.Printf("New Exchange Rate: %+v\n", exchangeRate)
}
```

### Converting CurrencyInfo to Map

You can convert a `CurrencyInfo` object to a map representation using the `ToMap` method.

```go
currencyMap := exchangeRate.ToMap()
fmt.Printf("CurrencyInfo as Map: %+v\n", currencyMap)
```

### Converting Map to CurrencyInfo

You can convert a map representation back to a `CurrencyInfo` entity using the `MapToCurrencyInfoEntity` function.

```go
currencyEntity, err := entity.MapToCurrencyInfoEntity(currencyMap)
if err != nil {
    log.Fatalf("Error converting map to CurrencyInfo: %v", err)
}

fmt.Printf("CurrencyInfo Entity: %+v\n", currencyEntity)
```

### Repository Interface

The `ExchangeRateRepositoryInterface` defines the methods that any repository implementation of `CurrencyInfo` must implement.

#### Methods

- `Save(currencyInfo *CurrencyInfo) error`: Saves a `CurrencyInfo` entity.
- `FindAll() ([]*CurrencyInfo, error)`: Retrieves all `CurrencyInfo` entities.
- `Find(code string, codeIn string) ([]*CurrencyInfo, error)`: Finds `CurrencyInfo` entities by currency codes.
- `FindByID(id string) (*CurrencyInfo, error)`: Finds a `CurrencyInfo` entity by its ID.
- `Delete(id string) error`: Deletes a `CurrencyInfo` entity by its ID.
