package evm

import (
	"context"
	"crypto/ecdsa"

	"github.com/CustodyOne/chainkit/signer"
	"github.com/CustodyOne/chainkit/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type LocalSigner struct {
	key *ecdsa.PrivateKey
}

func NewLocalSigner(key *ecdsa.PrivateKey) signer.Signer {
	return &LocalSigner{key}
}

func (s *LocalSigner) PublicKey(ctx context.Context) ([]byte, error) {
	pubkey := s.key.Public().(*ecdsa.PublicKey)
	return crypto.FromECDSAPub(pubkey), nil
}

func (s *LocalSigner) SharedKey(theirKey []byte) ([]byte, error) {
	return nil, nil
}

func (s *LocalSigner) Sign(payload types.TxDataToSign) (types.TxSignature, error) {
	return crypto.Sign(payload, s.key)
}
