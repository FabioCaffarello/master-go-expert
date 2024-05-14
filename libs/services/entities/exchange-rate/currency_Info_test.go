package exchangerateentity

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CurrencyInfoEntityTestSuite struct {
	suite.Suite
}

func TestCurrencyInfoEntityTestSuite(t *testing.T) {
	suite.Run(t, new(CurrencyInfoEntityTestSuite))
}
