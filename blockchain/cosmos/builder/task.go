package builder

import (
	"fmt"
	"math/big"
	"strings"

	"cosmossdk.io/math"
	"github.com/CustodyOne/chainkit/blockchain/cosmos/tx"
	"github.com/CustodyOne/chainkit/blockchain/cosmos/tx_input"
	xcbuilder "github.com/CustodyOne/chainkit/builder"
	xc "github.com/CustodyOne/chainkit/types"
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ xcbuilder.TxXTransferBuilder = &TxBuilder{}

func (txBuilder TxBuilder) NewTask(args *xcbuilder.TransferArgs, input xc.TxInput) (xc.Tx, error) {
	txInput := input.(*tx_input.TxInput)

	assetI, _ := args.GetAsset()
	asset := assetI.(*xc.TaskConfig)
	amountInt := big.Int(args.GetAmount())
	amountCoin := types.Coin{
		Denom:  txBuilder.GetDenom(asset),
		Amount: math.NewIntFromBigInt(&amountInt),
	}

	if strings.HasPrefix(asset.Code, "CosmosUndelegateOperator") {
		validatorAddress, ok := asset.DefaultParams["validator_address"]
		if !ok {
			return &tx.Tx{}, fmt.Errorf("must provide validator_address in task '%s'", asset.ID())
		}
		msgUndelegate := &stakingtypes.MsgUndelegate{
			DelegatorAddress: string(args.GetFrom()),
			Amount:           amountCoin,
			ValidatorAddress: fmt.Sprintf("%s", validatorAddress),
		}

		fees := txBuilder.calculateFees(asset, args.GetAmount(), txInput, false)
		return txBuilder.createTxWithMsg(txInput, msgUndelegate, txArgs{
			Memo:          txInput.LegacyMemo,
			FromPublicKey: txInput.LegacyFromPublicKey,
		}, fees)
	}

	return &tx.Tx{}, fmt.Errorf("not implemented task: '%s'", asset.ID())
}
