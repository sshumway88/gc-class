package main

import (
	"fmt"
	"log"

	"github.com/ardanlabs/blockchain/foundation/blockchain/database"
	"github.com/ardanlabs/blockchain/foundation/blockchain/signature"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	err := sendTx()

	if err != nil {
		log.Fatalln(err)
	}
}

func sendTx() error {
	privateKey, err := crypto.LoadECDSA("zblock/accounts/kennedy.ecdsa")
	if err != nil {
		return fmt.Errorf("unable to load private key for node: %w", err)
	}

	tx, err := database.NewTx(1, 1, "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32", "heather", 100, 0, nil)
	if err != nil {
		return fmt.Errorf("new tx: %w", err)
	}

	signedTx, err := tx.Sign(privateKey)
	if err != nil {
		return fmt.Errorf("sign tx: %w", err)
	}

	// =========================================================================

	tx2, err := database.NewTx(1, 1, "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32", "heather", 100, 0, nil)
	if err != nil {
		return fmt.Errorf("new tx2: %w", err)
	}

	addr, err := signature.FromAddress(tx2, signedTx.V, signedTx.R, signedTx.S)
	if err != nil {
		return fmt.Errorf("address: %w", err)
	}

	fmt.Println("*****>  0xF01813E4B85e178A83e29B8E7bF26BD830a25f32")
	fmt.Println("*****> ", tx2.FromID == addr)

	// data, err := json.Marshal(signedTx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := http.Post(fmt.Sprintf("%s/v1/tx/submit", "http://0.0.0.0:8080"), "application/json", bytes.NewBuffer(data))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	return nil
}
