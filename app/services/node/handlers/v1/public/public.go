// Package public maintains the group of handlers for public access.
package public

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	v1 "github.com/ardanlabs/blockchain/business/web/v1"
	"github.com/ardanlabs/blockchain/foundation/blockchain/database"
	"github.com/ardanlabs/blockchain/foundation/blockchain/signature"
	"github.com/ardanlabs/blockchain/foundation/web"
	"go.uber.org/zap"
)

// Handlers manages the set of bar ledger endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// SubmitWalletTransaction adds new transactions to the mempool.
func (h Handlers) SubmitWalletTransaction(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("web value missing from context")
	}

	// Decode the JSON in the post call into a Signed transaction.
	var signedTx database.SignedTx
	if err := web.Decode(r, &signedTx); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	if err := signature.VerifySignature(signedTx.V, signedTx.R, signedTx.S); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	addr, err := signature.FromAddress(signedTx.Tx, signedTx.V, signedTx.R, signedTx.S)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	h.Log.Infow("add tran", "traceid", v.TraceID, "sig:nonce", signedTx, "from", signedTx.FromID, "sig", addr, "to", signedTx.ToID, "value", signedTx.Value, "tip", signedTx.Tip)

	if addr != signedTx.FromID {
		return v1.NewRequestError(errors.New("sig not match"), http.StatusBadRequest)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// Sample just provides a starting point for the class.
func (h Handlers) Sample(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resp := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}
