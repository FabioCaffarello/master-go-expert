package gouuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoUuidTestSuite struct {
	suite.Suite
}

func TestGoUuidSuite(t *testing.T) {
	suite.Run(t, new(GoUuidTestSuite))
}

func (suite *GoUuidTestSuite) TestGetID() {
	properties := map[string]interface{}{
		"active":    true,
		"service":   "service",
		"source":    "source",
	}

	configID, err := GetID(properties)

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), configID)
	assert.Equal(suite.T(), "cd8282e9-c51f-5ae0-be5e-ea32024f3373", configID)
}


