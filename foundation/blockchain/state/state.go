// Package state is the core API for the blockchain and implements all the
// business rules and processing.
package state

import "github.com/ardanlabs/blockchain/foundation/blockchain/genesis"

// State manages the blockchain database.
type State struct {
	genesis genesis.Genesis
}

// Config represents the configuration required to start
// the blockchain node.
type Config struct {
	Genesis genesis.Genesis
}

// New constructs a new blockchain for data management.
func New(cfg Config) (*State, error) {

	// Create the State to provide support for managing the blockchain.
	state := State{
		genesis: cfg.Genesis,
	}

	return &state, nil
}
