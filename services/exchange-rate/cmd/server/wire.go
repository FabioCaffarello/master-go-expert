// go:build wireinject
// +build wireinject

package main

import (
	"libs/services/entities/exchange-rate/entity"
	"libs/services/infrastructure/database/repositories/exchange-rate/in-memory/go-doc-db/repository"
	webHandler "libs/services/infrastructure/server/http/webserver/handlers"
	inMemoryDBClient "libs/resources/database/in-memory/go-doc-db-client/client"
	"github.com/google/wire"
)

var setExchangeRateRepositoryDependency = wire.NewSet(
	repository.NewExchangeRateRepository,
	wire.bind(
		new(entity.ExchangeRateReositoryInterface),
		new(repository.ExchangeRateRepository),
	)
)

func NewWebServiceExchangeRateHandler(client *inMemoryDBClient.Client, databaseName string) *webHandler.ExchangeRateHandler {
	wire.Build(
		setExchangeRateRepositoryDependency,
		webHandler.NewWebServiceExchangeRateHandler,
	)
	return &webHandler.ExchangeRateHandler{}
}
