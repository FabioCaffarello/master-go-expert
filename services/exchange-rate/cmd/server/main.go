package main

import (
	inMemoryDBClient "libs/resources/database/in-memory/go-doc-db-client/client"
	inMemoryDB "libs/resources/database/in-memory/go-doc-db/database"
	webHandler "libs/services/infrastructure/server/http/handlers/exchange-rate"
	"libs/services/infrastructure/server/http/webserver"
	"log"
	"net/http"
)

var (
	dbName        = "exchange-rate"
	webserverPort = ":8080"
)

func RegisterExchangeRateWebServerTransportRoutes(server *webserver.Server, webService *webHandler.WebServiceExchangeRateHandler) {
	server.RegisterRoute(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Home Page!"))
	})

	server.RegisterRoute(http.MethodGet, "/cotacoes", webService.ListCurrentExchangeRate)
}

func main() {
	db := inMemoryDB.NewInMemoryDocBD(dbName)
	dbClient := inMemoryDBClient.NewClient(db)

	webserver := webserver.NewWebServer(webserverPort)
	webserver.ConfigureDefaults()
	webServiceExchangeRate := NewWebServiceExchangeRateHandler(dbClient, dbName)
	RegisterExchangeRateWebServerTransportRoutes(webserver, webServiceExchangeRate)

	if err := webserver.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
