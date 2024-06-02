package handlers

import (
	"libs/services/api-clients/exchange-rate/client"
	gosd "libs/shared/go-sd/service_discovery"
	"log"
	"net/http"
)

var (
	code   = "USD"
	codeIn = "BRL"
)

type WebCotacoesHandler struct {
	client *client.Client
}

func NewWebCotacoesHandler(sd *gosd.ServiceDiscovery) *WebCotacoesHandler {
	return &WebCotacoesHandler{
		client: client.NewClient(sd),
	}
}

func (h *WebCotacoesHandler) HandleGetCotacoes(w http.ResponseWriter, r *http.Request) error {
	currencyInfo, err := h.client.GetExchangeRate(code, codeIn)
	log.Printf("currencyInfo: %v", currencyInfo)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return WriteJSON(w, http.StatusOK, currencyInfo)
}
