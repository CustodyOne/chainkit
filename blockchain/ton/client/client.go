package client

import (
	"github.com/CustodyOne/chainkit/blockchain/ton/client/liteserver"
	"github.com/CustodyOne/chainkit/blockchain/ton/client/tonapi"
	"github.com/CustodyOne/chainkit/client"
	xc "github.com/CustodyOne/chainkit/types"
)

type TonApiProvider string

var TonApi TonApiProvider = "tonapi"
var LiteServer TonApiProvider = "liteserver"

type TonClient interface {
	client.IClient
}

func NewClient(cfg *xc.ChainConfig) (TonClient, error) {
	switch TonApiProvider(cfg.Client.Provider) {
	case TonApi:
		return tonapi.NewClient(cfg)
	case LiteServer:
		return liteserver.NewClient(cfg)
	default:
		return tonapi.NewClient(cfg)
	}
}
