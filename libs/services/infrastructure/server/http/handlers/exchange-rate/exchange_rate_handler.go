package handlers

import (
	"encoding/json"
	entity "libs/services/entities/exchange-rate/entity"
	usecase "libs/services/usecases/exchange-rate/usecases"
	"net/http"
)

type WebServiceExchangeRateHandler struct {
	ExchangeRateRepository entity.ExchangeRateReositoryInterface
}

func NewWebServiceExchangeRateHandler(
	exchangeRateRepository entity.ExchangeRateReositoryInterface,
) *WebServiceExchangeRateHandler {
	return &WebServiceExchangeRateHandler{
		ExchangeRateRepository: exchangeRateRepository,
	}
}

func (h *WebServiceExchangeRateHandler) ListCurrentExchangeRate(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	codeIn := r.URL.Query().Get("code_in")
	if code == "" || codeIn == "" {
		code = "USD"
		codeIn = "BRL"
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
