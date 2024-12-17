package testutil

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// NewInMemoryDB returns an in-memory database for testing
func NewInMemoryDB() dbm.DB {
	return dbm.NewMemDB()
}

// NewContext creates a new context for testing
func NewContext(db dbm.DB) sdk.Context {
	cms := store.NewCommitMultiStore(db)
	return sdk.NewContext(cms, false, nil)
}

// NewTestCodec returns a codec for testing
func NewTestCodec() codec.Codec {
	interfaceRegistry := codecTypes.NewInterfaceRegistry()
	return codec.NewProtoCodec(interfaceRegistry)
}

// NewTestStoreKey creates a store key for testing
func NewTestStoreKey(name string) sdk.StoreKey {
	return sdk.NewKVStoreKey(name)
}

// MockTx implements a mock transaction for testing
type MockTx struct {
	data []byte
}

func NewMockTx(data []byte) MockTx {
	return MockTx{data: data}
}

func (tx MockTx) GetMsgs() []sdk.Msg {
	return nil
}

// Mock bank keeper for testing
type MockBankKeeper struct{}

func NewMockBankKeeper() MockBankKeeper {
	return MockBankKeeper{}
}

// NoopAnteHandler returns a do-nothing ante handler for testing
func NoopAnteHandler() sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		return ctx, nil
	}
}

// GenerateAddress generates a random address for testing
func GenerateAddress() sdk.AccAddress {
	return sdk.AccAddress([]byte("test_address"))
}

func CreateTestContext(store store.CommitMultiStore) sdk.Context {
	ms := store.(sdk.MultiStore)
	header := tmproto.Header{
		Height: 1,
		Time:   time.Now(),
	}
	logger := tmlog.NewNopLogger()
	return sdk.NewContext(ms, header, false, logger)
}
