package evm_legacy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChainkitTestSuite struct {
	suite.Suite
	Ctx context.Context
}

func (s *ChainkitTestSuite) SetupTest() {
	s.Ctx = context.Background()
}

func TestLegacyEvm(t *testing.T) {
	suite.Run(t, new(ChainkitTestSuite))
}
