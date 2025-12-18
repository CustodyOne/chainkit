package btc_cash

import (
	"github.com/CustodyOne/chainkit/blockchain/btc/client"
	xc_client "github.com/CustodyOne/chainkit/client"
	xc "github.com/CustodyOne/chainkit/types"
)

func NewClient(cfg *xc.ChainConfig) (xc_client.IClient, error) {
	cli, err := client.NewBitcoinClient(cfg)
	if err != nil {
		return cli, err
	}
	return cli.WithAddressDecoder(&BchAddressDecoder{}).(client.BtcClient), nil
}
