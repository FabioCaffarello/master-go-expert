# Exchange Rate Use Cases

The `exchange-rate` library provides use cases for handling exchange rate operations, including fetching and saving exchange rates.

## Overview

This package includes the following main components:
- `GetExchangeRateUseCase`: A struct that provides methods to fetch exchange rates from an external API, save them to a repository, and return the exchange rate data.
- `GenerateExchangeRateSearchKey`: A function that generates a search key for the exchange rate by concatenating and uppercasing the currency codes.

## Features

The main functionalities provided by the package include:
- Fetching exchange rates from an external API.
- Saving exchange rates to a repository.
- Generating a search key for exchange rates.

## Types

- **GetExchangeRateUseCase**: Represents a use case for fetching and saving exchange rates.

## Functions

### GetExchangeRateUseCase Functions

- `NewGetExchangeRateUseCase(repository entity.ExchangeRateRepositoryInterface) *GetExchangeRateUseCase`: Creates and returns a new `GetExchangeRateUseCase` instance.
- `Execute(code, codeIn string) (outputDTO.ExchangeRatesDTO, error)`: Fetches the exchange rate for the given currency codes, saves it to the repository, and returns the exchange rate data.

### Utility Functions

- `GenerateExchangeRateSearchKey(code, codeIn string) string`: Generates a search key for the exchange rate by concatenating and uppercasing the currency codes.

## Usage

### Creating a New Use Case

```go
exchangeRateRepository := // initialize your ExchangeRateRepository
useCase := usecases.NewGetExchangeRateUseCase(exchangeRateRepository)
```

### Executing the Use Case

```go
code := "USD"
codeIn := "BRL"

exchangeRates, err := useCase.Execute(code, codeIn)
if err != nil {
    log.Fatal(err)
}

fmt.Println(exchangeRates)
```

## Full Example
Here is a complete example demonstrating the setup and use of the usecases package in a simple application:

```go
package main

import (
	"fmt"
	"log"
	"libs/external-clients/economia-awesome-api/client"
	"libs/resources/database/in-memory/go-doc-db-client/client"
	"libs/resources/database/in-memory/go-doc-db/database"
	"libs/services/acl/dtos/exchange-rate/output"
	"libs/services/entities/exchange-rate/entity"
	"libs/services/usecases/exchange-rate/usecases"
)

func main() {
	// Initialize the in-memory database and repository
	dbClient := client.NewClient(database.NewInMemoryDocBD("myDatabase"))
	exchangeRateRepository := entity.NewExchangeRateRepository(dbClient)

	// Create the use case
	useCase := usecases.NewGetExchangeRateUseCase(exchangeRateRepository)

	// Execute the use case to get exchange rates
	code := "USD"
	codeIn := "BRL"

	exchangeRates, err := useCase.Execute(code, codeIn)
	if err != nil {
		log.Fatal(err)
	}

	// Print the exchange rates
	fmt.Println(exchangeRates)
}
```

In this example, we initialize an in-memory database, create an exchange rate repository, set up the use case for fetching exchange rates, and execute the use case to retrieve and print the exchange rates.
