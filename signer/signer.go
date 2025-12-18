package signer

import (
	"context"

	"github.com/CustodyOne/chainkit/types"
)

type Signer interface {
	PublicKey(ctx context.Context) ([]byte, error)
	SharedKey(theirKey []byte) ([]byte, error)
	Sign(payload types.TxDataToSign) (types.TxSignature, error)
}
