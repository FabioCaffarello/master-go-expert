package usecases

import (
	"libs/external-clients/economia-awesome-api/client"
	outputDTO "libs/services/acl/dtos/exchange-rate/output"
	entity "libs/services/entities/exchange-rate/entity"
)

type GetExchangeRateUseCase struct {
	repository               entity.ExchangeRateReositoryInterface
	economiaAwesomeApiClient *client.Client
}

func NewGetExchangeRateUseCase(
	repository entity.ExchangeRateReositoryInterface,
) *GetExchangeRateUseCase {
	return &GetExchangeRateUseCase{
		repository:               repository,
		economiaAwesomeApiClient: client.NewClient(),
	}
}

func (u *GetExchangeRateUseCase) Execute(code, codeIn string) (outputDTO.ExchangeRatesDTO, error) {
	searchKey := GenerateExchangeRateSearchKey(code, codeIn)

	awesomeAPIresult, err := u.economiaAwesomeApiClient.GetExchangeRate(searchKey)
	if err != nil {
		return outputDTO.ExchangeRatesDTO{}, err
	}

	output := make(outputDTO.ExchangeRatesDTO, 0)
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
