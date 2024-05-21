package exchangerateentity

// ExchangeRateRepositoryInterface defines the methods that any repository implementation
// of CurrencyInfo must implement.
type ExchangeRateRepositoryInterface interface {
	Save(currencyInfo *CurrencyInfo) error
	FindAll() ([]*CurrencyInfo, error)
	Find(code string, codeIn string) ([]*CurrencyInfo, error)
	FindByID(id string) (*CurrencyInfo, error)
	Delete(id string) error
}
