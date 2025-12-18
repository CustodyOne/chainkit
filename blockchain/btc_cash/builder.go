package btc_cash

import (
	"github.com/CustodyOne/chainkit/blockchain/btc"
	"github.com/CustodyOne/chainkit/blockchain/btc/tx"
	xcbuilder "github.com/CustodyOne/chainkit/builder"
	xc "github.com/CustodyOne/chainkit/types"
)

// TxBuilder for Bitcoin
type TxBuilder struct {
	btc.TxBuilder
}

var _ xcbuilder.TxBuilder = &TxBuilder{}

// NewTxBuilder creates a new Bitcoin TxBuilder
func NewTxBuilder(cfg *xc.ChainConfig) (TxBuilder, error) {
	txBuilder, err := btc.NewTxBuilder(cfg)
	if err != nil {
		return TxBuilder{}, err
	}
	return TxBuilder{
		TxBuilder: txBuilder.WithAddressDecoder(&BchAddressDecoder{}),
	}, nil
}

func (txBuilder TxBuilder) NewTransfer(args *xcbuilder.TransferArgs, input xc.TxInput) (xc.Tx, error) {
	txObj, err := txBuilder.TxBuilder.NewTransfer(args, input)
	if err != nil {
		return txObj, err
	}
	return txObj.(*tx.Tx), nil
}

func (txBuilder TxBuilder) NewNativeTransfer(args *xcbuilder.TransferArgs, input xc.TxInput) (xc.Tx, error) {
	txObj, err := txBuilder.TxBuilder.NewNativeTransfer(args, input)
	if err != nil {
		return txObj, err
	}
	return txObj.(*tx.Tx), nil
}

func (txBuilder TxBuilder) NewTokenTransfer(args *xcbuilder.TransferArgs, input xc.TxInput) (xc.Tx, error) {
	txObj, err := txBuilder.TxBuilder.NewTokenTransfer(args, input)
	if err != nil {
		return txObj, err
	}
	return txObj.(*tx.Tx), nil
}
