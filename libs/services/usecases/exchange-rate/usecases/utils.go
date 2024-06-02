package usecases

import "strings"

// GenerateExchangeRateSearchKey generates a search key for the exchange rate by concatenating and uppercasing the currency codes.
func GenerateExchangeRateSearchKey(code, codeIn string) string {
	return strings.ToUpper(code + "-" + codeIn)
}
