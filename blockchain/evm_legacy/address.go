package evm_legacy

import (
	evmaddress "github.com/CustodyOne/chainkit/blockchain/evm/address"
	xc "github.com/CustodyOne/chainkit/types"
)

type AddressBuilder = evmaddress.AddressBuilder

var NewAddressBuilder = evmaddress.NewAddressBuilder

var _ xc.AddressBuilder = AddressBuilder{}
