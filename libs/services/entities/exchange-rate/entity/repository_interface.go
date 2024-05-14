package exchangerateentity

type ExchangeRateReositoryInterface interface {
	FindAll() ([]*CurrencyInfo, error)
	FindByID(id string) (*CurrencyInfo, error)
	Save(exchangeRate *CurrencyInfo) error
}
