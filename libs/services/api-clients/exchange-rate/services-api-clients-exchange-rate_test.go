package exchangerate

import (
	"testing"
)

func TestServicesApiClientsExchangeRate(t *testing.T) {
	result := ServicesApiClientsExchangeRate("works")
	if result != "ServicesApiClientsExchangeRate works" {
		t.Error("Expected ServicesApiClientsExchangeRate to append 'works'")
	}
}
