package app

import (
	"reflect"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"

	"github.com/sei-protocol/sei-chain/app/antedecorators"
	mevkeeper "github.com/sei-protocol/sei-chain/x/mev/keeper"
)

// HandlerOptions are the options required for constructing an AnteHandler.
type HandlerOptions struct {
	AccountKeeper     authkeeper.AccountKeeper
	BankKeeper        bankkeeper.Keeper
	IBCKeeper         *ibckeeper.Keeper
	ParamsKeeper      paramskeeper.Keeper
	SignModeHandler   signing.SignModeHandler
	FeegrantKeeper    ante.FeegrantKeeper
	SigGasConsumer    ante.SignatureVerificationGasConsumer
	MevKeeper         mevkeeper.Keeper
	TxConfig          client.TxConfig
	TXCounterStoreKey types.StoreKey
	WasmConfig        *wasmtypes.WasmConfig
	WasmKeeper        *wasmkeeper.Keeper
	OracleKeeper      *oraclekeeper.Keeper
	EVMKeeper         *evmkeeper.Keeper
	TracingInfo       *types.TraceContext
	LatestCtxGetter   func() sdk.Context
}

// ChainDecorators chains multiple decorators
func ChainDecorators(decorators ...sdk.AnteDecorator) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
		for _, decorator := range decorators {
			newCtx, err = decorator.AnteHandle(ctx, tx, simulate, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
				return ctx, nil
			})
			if err != nil {
				return newCtx, err
			}
			ctx = newCtx
		}
		return ctx, nil
	}
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(opts HandlerOptions) (sdk.AnteHandler, error) {
	if reflect.ValueOf(opts.AccountKeeper).IsZero() ||
		reflect.ValueOf(opts.BankKeeper).IsZero() ||
		reflect.ValueOf(opts.SignModeHandler).IsZero() ||
		reflect.ValueOf(opts.MevKeeper).IsZero() ||
		opts.TxConfig == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "all keepers and handlers are required")
	}

	decorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(func(isCheckTx bool, ctx sdk.Context, txBytes uint64, tx sdk.Tx) sdk.Context {
			return ctx
		}),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(opts.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(opts.AccountKeeper),
		ante.NewDeductFeeDecorator(opts.AccountKeeper, opts.BankKeeper, opts.FeegrantKeeper, opts.ParamsKeeper, nil),
		ante.NewSetPubKeyDecorator(opts.AccountKeeper),
		ante.NewValidateSigCountDecorator(opts.AccountKeeper),
		ante.NewSigGasConsumeDecorator(opts.AccountKeeper, opts.SigGasConsumer),
		ante.NewSigVerificationDecorator(opts.AccountKeeper, opts.SignModeHandler),
		ante.NewIncrementSequenceDecorator(opts.AccountKeeper),
		antedecorators.NewPriorityDecorator(opts.MevKeeper, opts.TxConfig),
	}

	return ChainDecorators(decorators...), nil
}

func NewAnteHandlerAndDepGenerator(opts HandlerOptions) (sdk.AnteHandler, sdk.AnteDepGenerator, error) {
	handler, err := NewAnteHandler(opts)
	if err != nil {
		return nil, nil, err
	}
	return handler, nil, nil
}
