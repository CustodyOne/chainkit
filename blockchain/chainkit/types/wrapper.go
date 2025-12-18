package types

import (
	xc_types "github.com/CustodyOne/chainkit/types"
)

type TxInputWrapper struct {
	xc_types.TxInputEnvelope
	xc_types.TxInput
}
