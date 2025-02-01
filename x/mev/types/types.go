package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// Bundle represents a MEV bundle submission
// Defined in the protos.
// type Bundle struct {
// 	Sender    sdk.AccAddress `json:"sender"`
// 	Txs       []string       `json:"txs"`        // List of encoded transactions
// 	BlockNum  uint64         `json:"block_num"`  // Target block number
// 	Timestamp int64         `json:"timestamp"`
// }

// BundleResponse is the response type for bundle queries
// type BundleResponse struct {
// 	Bundles []Bundle `json:"bundles"`
// }

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitBundle{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

//
//func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
//	//registry.RegisterImplementations((*govtypes.Content)(nil),
//	//	&AddERCNativePointerProposal{},
//	//	&AddERCCW20PointerProposal{},
//	//	&AddERCCW721PointerProposal{},
//	//	&AddERCCW1155PointerProposal{},
//	//	&AddCWERC20PointerProposal{},
//	//	&AddCWERC721PointerProposal{},
//	//	&AddCWERC1155PointerProposal{},
//	//	&AddERCNativePointerProposalV2{},
//	//)
//	//registry.RegisterImplementations(
//	//	(*sdk.Msg)(nil),
//	//	&MsgEVMTransaction{},
//	//	&MsgSend{},
//	//	&MsgRegisterPointer{},
//	//	&MsgAssociateContractAddress{},
//	//)
//	//registry.RegisterInterface(
//	//	"seiprotocol.seichain.evm.TxData",
//	//	(*ethtx.TxData)(nil),
//	//	&ethtx.DynamicFeeTx{},
//	//	&ethtx.AccessListTx{},
//	//	&ethtx.LegacyTx{},
//	//	&ethtx.BlobTx{},
//	//	&ethtx.AssociateTx{},
//	//)
//
//	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
//}
