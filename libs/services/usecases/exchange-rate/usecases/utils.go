package usecases

import "strings"

func GenerateExchangeRateSearchKey(code, codeIn string) string {
	return strings.ToUpper(code + codeIn)
}
