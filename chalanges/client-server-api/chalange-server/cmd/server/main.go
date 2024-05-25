package main

import (
	webhandlers "chalanges/client-server-api/chalange-server/internal/handlers"
	"flag"
	gosd "libs/shared/go-sd/service_discovery"
	"log"
	"net/http"
)

var (
	webServerPort  = ":5000"
	useHostAddress = false
)

func main() {
	listenAddr := flag.String("listenAddr", webServerPort, "The address to listen on for HTTP requests.")
	flag.Parse()

	sd := gosd.NewServiceDiscovery(useHostAddress)
	cotacoesHandler := webhandlers.NewWebCotacoesHandler(sd)

	http.HandleFunc("/cotacoes", webhandlers.MakeAPIFunc(cotacoesHandler.HandleGetCotacoes))
	log.Printf("Server listening on %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
