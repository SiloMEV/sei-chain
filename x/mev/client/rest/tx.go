package rest

import (
	"encoding/base64"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/sei-protocol/sei-chain/x/mev/types"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/mev/bundles", submitBundleHandler(clientCtx)).Methods("POST")
}

type submitBundleRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Transactions []string     `json:"transactions"` // base64 encoded transactions
	BlockHeight  int64        `json:"block_height"`
	BundleFee    string       `json:"bundle_fee"`
}

func submitBundleHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req submitBundleRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		bundleFee, err := sdk.ParseCoinNormalized(req.BundleFee)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Convert base64 transactions to bytes
		txs := make([][]byte, len(req.Transactions))
		for i, txStr := range req.Transactions {
			tx, err := base64.StdEncoding.DecodeString(txStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			txs[i] = tx
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSubmitBundle(
			fromAddr,
			txs,
			req.BlockHeight,
			bundleFee,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
