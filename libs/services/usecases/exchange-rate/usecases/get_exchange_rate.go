package usecases

import (
	"errors"
	"libs/external-clients/economia-awesome-api/client"
	outputDTO "libs/services/acl/dtos/exchange-rate/output"
	entity "libs/services/entities/exchange-rate/entity"
	"log"
)

// GetExchangeRateUseCase represents a use case for fetching and saving exchange rates.
type GetExchangeRateUseCase struct {
	repository               entity.ExchangeRateRepositoryInterface
	economiaAwesomeApiClient *client.Client
}

// NewGetExchangeRateUseCase creates and returns a new GetExchangeRateUseCase instance.
func NewGetExchangeRateUseCase(
	repository entity.ExchangeRateRepositoryInterface,
) *GetExchangeRateUseCase {
	return &GetExchangeRateUseCase{
		repository:               repository,
		economiaAwesomeApiClient: client.NewClient(),
	}
}

// Execute fetches the exchange rate for the given currency codes, saves it to the repository, and returns the exchange rate data.
func (u *GetExchangeRateUseCase) Execute(code, codeIn string) (outputDTO.ExchangeRatesDTO, error) {
	if code == "" || codeIn == "" {
		return outputDTO.ExchangeRatesDTO{}, errors.New("currency codes cannot be empty")
	}
	log.Printf("Getting exchange rate from Economia Awesome API for %s/%s", code, codeIn)
	searchKey := GenerateExchangeRateSearchKey(code, codeIn)
	log.Printf("Search key: %s", searchKey)

	awesomeAPIresult, err := u.economiaAwesomeApiClient.GetExchangeRate(searchKey)
	log.Printf("API result: %v", awesomeAPIresult)
	if err != nil {
		return outputDTO.ExchangeRatesDTO{}, err
	}

	output := make(outputDTO.ExchangeRatesDTO)
	for _, rate := range awesomeAPIresult {
		exchangeRate, err := entity.NewExchangeRate(
			rate.Code,
			rate.CodeIn,
			rate.Name,
			rate.High,
			rate.Low,
			rate.VarBid,
			rate.PctChange,
			rate.Bid,
			rate.Ask,
			rate.Timestamp,
			rate.CreateDate,
		)
		if err != nil {
			return outputDTO.ExchangeRatesDTO{}, err
		}
		err = u.repository.Save(exchangeRate)
		if err != nil {
			return outputDTO.ExchangeRatesDTO{}, err
		}
		exchangeRateOut := outputDTO.ExchangeRateDTO{
			Code:       exchangeRate.Code,
			CodeIn:     exchangeRate.CodeIn,
			Name:       exchangeRate.Name,
			High:       exchangeRate.High,
			Low:        exchangeRate.Low,
			VarBid:     exchangeRate.VarBid,
			PctChange:  exchangeRate.PctChange,
			Bid:        exchangeRate.Bid,
			Ask:        exchangeRate.Ask,
			Timestamp:  exchangeRate.Timestamp,
			CreateDate: exchangeRate.CreateDate.Format("2006-01-02 15:04:05"),
		}
		output[searchKey] = exchangeRateOut
	}

	return output, nil
}
