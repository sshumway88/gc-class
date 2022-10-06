// Package private maintains the group of handlers for node to node access.
package private

import (
	"context"
	"net/http"

	"github.com/ardanlabs/blockchain/foundation/blockchain/state"
	"github.com/ardanlabs/blockchain/foundation/web"
	"go.uber.org/zap"
)

// Handlers manages the set of bar ledger endpoints.
type Handlers struct {
	Log   *zap.SugaredLogger
	State *state.State
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
