package evm_legacy

import (
	evmtx "github.com/CustodyOne/chainkit/blockchain/evm/tx"
	xc "github.com/CustodyOne/chainkit/types"
)

// Tx for EVM
type Tx = evmtx.Tx

var _ xc.Tx = &Tx{}
