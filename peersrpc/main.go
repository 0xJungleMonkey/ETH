package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/p2p"
)


func main() {
	// define the post properties
	url := "http://localhost:8545"

	payload := `{
	"jsonrpc": "2.0",
	"method": "admin_peers",
	"params": [],
	"id": 1
}`
//http post
	resp, err := http.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		log.Fatal("Failed to connect to Geth client:", err)
	}
	defer resp.Body.Close()

	var result struct {
		Result []*p2p.PeerInfo `json:"result"`
		Error  struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
//Read response into result struct 
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal("Failed to parse response:", err)
	}

	if result.Error.Message != "" {
		log.Fatal("RPC error:", result.Error.Message)
	}

	log.Println("Connected Peers:")
	// print results
	for _, peer := range result.Result {
		log.Printf("ID: %s, Name: %s, Local Address: %s, Remote Address: %s\n", peer.ID, peer.Name, peer.Network.LocalAddress, peer.Network.RemoteAddress)
	}
}
