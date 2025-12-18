package protocols_test

import (
	"testing"

	"github.com/CustodyOne/chainkit/factory/protocols"
	xc "github.com/CustodyOne/chainkit/types"
	"github.com/stretchr/testify/suite"
)

type BlockchainTestSuite struct {
	suite.Suite
}

func TestBlockchainTestSuite(t *testing.T) {
	suite.Run(t, new(BlockchainTestSuite))
}

func (suite *BlockchainTestSuite) SetupTest(t *testing.T) {
}

func (s *BlockchainTestSuite) TestAllNewClient() {
	require := s.Require()

	for _, protocol := range xc.SupportedProtocols {
		/*
			// TODO: these require custom params for NewClient
			if protocol == xc.ProtocolAptos || protocol == xc.ProtocolSubstrate {
				continue
			}
		*/

		res, err := protocols.NewClient(createChainFor(protocol), protocol)
		require.NoError(err, "Missing protocol for NewClient: "+protocol)
		require.NotNil(res)
	}
}

func createChainFor(protocol xc.Protocol) *xc.ChainConfig {
	fakeAsset := &xc.ChainConfig{
		Client:   &xc.ClientConfig{},
		Protocol: protocol,
	}
	if protocol == xc.ProtocolBtc {
		fakeAsset.Chain = "BTC"
		fakeAsset.Client.Auth = "1234"
	}
	if protocol == xc.ProtocolBtcLegacy {
		fakeAsset.Chain = "DOGE"
		fakeAsset.Client.Auth = "1234"
	}
	if protocol == xc.ProtocolBtcCash {
		fakeAsset.Chain = "BCH"
		fakeAsset.Client.Auth = "1234"
	}
	/*
		if blockchain == xc.BlockchainSubstrate {
			fakeAsset.ChainPrefix = "0"
		}
	*/
	return fakeAsset
}
