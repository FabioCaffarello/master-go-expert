# Exchange Rate Handlers

The `exchange-rate` library provides HTTP handlers for exchange rate operations, utilizing the exchange rate use cases and entities.

## Overview

The `handlers` package includes the following main components:
- `WebServiceExchangeRateHandler`: A struct that provides methods to handle HTTP requests related to exchange rates.

## Features

The main functionalities provided by the package include:
- Handling HTTP GET requests to list the current exchange rate.

## Types

- **WebServiceExchangeRateHandler**: Handles HTTP requests for exchange rate operations.

## Functions

### WebServiceExchangeRateHandler Functions

- `NewWebServiceExchangeRateHandler(exchangeRateRepository entity.ExchangeRateRepositoryInterface) *WebServiceExchangeRateHandler`: Creates and returns a new `WebServiceExchangeRateHandler` instance.
- `ListCurrentExchangeRate(w http.ResponseWriter, r *http.Request)`: Handles HTTP GET requests to list the current exchange rate.

## Usage

### Creating a New Handler

```go
exchangeRateRepository := // initialize your ExchangeRateRepository
handler := handlers.NewWebServiceExchangeRateHandler(exchangeRateRepository)
```

### Handling HTTP Requests

```go
http.HandleFunc("/exchange-rate", handler.ListCurrentExchangeRate)
```

### Example

Here is a complete example of setting up the handler with an HTTP server:

```go
package main

import (
	"libs/resources/database/in-memory/go-doc-db-client/client"
	"libs/resources/database/in-memory/go-doc-db/database"
	"libs/services/entities/exchange-rate/entity"
	"libs/services/usecases/exchange-rate/usecases"
	"libs/web/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialize the in-memory database and repository
	dbClient := client.NewClient(database.NewInMemoryDocBD("myDatabase"))
	exchangeRateRepository := entity.NewExchangeRateRepository(dbClient)

	// Create the handler
	handler := handlers.NewWebServiceExchangeRateHandler(exchangeRateRepository)

	// Register the route and start the server
	http.HandleFunc("/exchange-rate", handler.ListCurrentExchangeRate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

In this example, we set up an in-memory database, initialize the repository, create the handler, and register an HTTP route to handle exchange rate requests.
