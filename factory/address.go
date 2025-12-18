package factory

import "github.com/CustodyOne/chainkit/types"

// MustAddress coverts a string to Address, panic if error
func (f *Factory) MustAddress(cfg types.IAsset, addressStr string) types.Address {
	return types.Address(addressStr)
}
