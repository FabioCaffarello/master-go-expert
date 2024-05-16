package main

import (
	inMemoryDBClient "libs/resources/database/in-memory/go-doc-db-client/client"
	inMemoryDB "libs/resources/database/in-memory/go-doc-db/database"
	"libs/services/infrastructure/server/http/webserver"
)

var (
	dbName = "exchange-rate"
	webserverPort = "8000"
)

func main() {
	db := inMemoryDB.NewInMemoryDocBD(dbName)
	dbClient := inMemoryDBClient.NewClient(db)

	webserver := webserver.NewWebServer(webserverPort)
}
