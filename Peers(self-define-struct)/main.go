package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

type PeerInfo struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Caps    []string `json:"caps"`
	Network struct {
		LocalAddress  string `json:"localAddress"`
		RemoteAddress string `json:"remoteAddress"`
	} `json:"network"`
	Enode string `json:"enode"`
}

func main() {
	// Connect to the local Geth client
	client, err := rpc.Dial("http://localhost:8545") // Change the URL if your Geth client is running on a different address

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}

	// Get the list of connected peers
	var peers []*PeerInfo
	err = client.Call(&peers, "admin_peers")

	if err != nil {
		log.Fatalf("Failed to retrieve the list of peers: %v\n", err)
	}

	// Display the list of peers
	fmt.Println("Connected Peers:")
	for i, peer := range peers {
		fmt.Printf("%d. ID: %s, Name: %s, Enode: %s\n", i+1, peer.ID, peer.Name, peer.Enode)
	}
}
