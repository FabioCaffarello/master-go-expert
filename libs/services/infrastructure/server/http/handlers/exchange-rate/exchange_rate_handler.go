package handlers

import (
	"encoding/json"
	"errors"
	entity "libs/services/entities/exchange-rate/entity"
	usecase "libs/services/usecases/exchange-rate/usecases"
	"net/http"
)

// WebServiceExchangeRateHandler handles HTTP requests for exchange rate operations.
type WebServiceExchangeRateHandler struct {
	ExchangeRateRepository entity.ExchangeRateRepositoryInterface
}

// NewWebServiceExchangeRateHandler creates and returns a new WebServiceExchangeRateHandler instance.
func NewWebServiceExchangeRateHandler(
	exchangeRateRepository entity.ExchangeRateRepositoryInterface,
) *WebServiceExchangeRateHandler {
	return &WebServiceExchangeRateHandler{
		ExchangeRateRepository: exchangeRateRepository,
	}
}

// ListCurrentExchangeRate handles HTTP GET requests to list the current exchange rate.
// It expects "code" and "code_in" query parameters, defaulting to "USD" and "BRL" respectively if not provided.
func (h *WebServiceExchangeRateHandler) ListCurrentExchangeRate(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	codeIn := r.URL.Query().Get("code_in")
	if code == "" || codeIn == "" {
		// code = "USD"
		// codeIn = "BRL"
		http.Error(w, errors.New("missing required query parameters 'code' and 'code_in").Error(), http.StatusBadRequest)
		return
	}

	getExchangeRate := usecase.NewGetExchangeRateUseCase(h.ExchangeRateRepository)

	exchangeRate, err := getExchangeRate.Execute(code, codeIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(exchangeRate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
