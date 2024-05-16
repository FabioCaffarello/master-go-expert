package exchangerateentity

type ExchangeRateRepositoryInterface interface {
	Save(exchangeRate *CurrencyInfo) error
	FindAll() ([]*CurrencyInfo, error)
	Find(code string, codeIn string) ([]*CurrencyInfo, error)
	FindByID(id string) (*CurrencyInfo, error)
}
