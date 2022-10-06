package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ardanlabs/blockchain/foundation/blockchain/database"
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

	data, err := json.Marshal(signedTx)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/v1/tx/submit", "http://0.0.0.0:8080"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("WE HAD A PROBLEM")
	}

	var result struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	return nil
}
