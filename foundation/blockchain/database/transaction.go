package database

// Tx is the transactional information between two parties.
type Tx struct {
	ChainID uint16 `json:"chain_id"` // Ethereum: The chain id that is listed in the genesis file.
	Nonce   uint64 `json:"nonce"`    // Ethereum: Unique id for the transaction supplied by the user.
	FromID  string `json:"from"`     // Ethereum: Account sending the transaction. Will be checked against signature.
	ToID    string `json:"to"`       // Ethereum: Account receiving the benefit of the transaction.
	Value   uint64 `json:"value"`    // Ethereum: Monetary value received from this transaction.
	Tip     uint64 `json:"tip"`      // Ethereum: Tip offered by the sender as an incentive to mine this transaction.
	Data    []byte `json:"data"`     // Ethereum: Extra data related to the transaction.
}

// NewTx constructs a new transaction.
func NewTx(chainID uint16, nonce uint64, fromID string, toID string, value uint64, tip uint64, data []byte) (Tx, error) {
	tx := Tx{
		ChainID: chainID,
		Nonce:   nonce,
		FromID:  fromID,
		ToID:    toID,
		Value:   value,
		Tip:     tip,
		Data:    data,
	}

	return tx, nil
}
