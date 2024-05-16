// go:build wireinject
// +build wireinject

package main

import ( 
	"github.com/google/wire"
	inMemoryDBClient "libs/resources/database/in-memory/go-doc-db-client/client"
	entity "libs/services/entities/exchange-rate/entity"
	repository "libs/services/infrastructure/database/repositories/exchange-rate/in-memory/go-doc-db/repository"
	webHandler "libs/services/infrastructure/server/http/handlers/exchange-rate"
)

var setExchangeRateRepositoryDependency = wire.NewSet(
	repository.NewExchangeRateRepository,
	wire.Bind(
		new(entity.ExchangeRateRepositoryInterface),
		new(*repository.ExchangeRateRepository),
	),
)

func NewWebServiceExchangeRateHandler(client *inMemoryDBClient.Client, databaseName string) *webHandler.WebServiceExchangeRateHandler {
	wire.Build(
		setExchangeRateRepositoryDependency,
		webHandler.NewWebServiceExchangeRateHandler,
	)
	return &webHandler.WebServiceExchangeRateHandler{}
}
