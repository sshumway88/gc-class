package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ardanlabs/blockchain/foundation/blockchain/database"
)

func main() {
	err := sendTx()

	if err != nil {
		log.Fatalln(err)
	}
}

func sendTx() error {
	tx, err := database.NewTx(1, 1, "bill", "heather", 100, 0, nil)
	if err != nil {
		return fmt.Errorf("new tx: %w", err)
	}

	data, err := json.Marshal(tx)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/v1/tx/submit", "http://0.0.0.0:8080"), "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return nil
}
