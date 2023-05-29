package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	//rpc.Dial() establish a connection to ETH client and return the rpc.Client
	client, err := rpc.Dial("http://localhost:8545") // Replace with the appropriate Ethereum client URL
	if err != nil {
		log.Fatal(err)
	}
	// Define array to get the rpc result 
	var peerInfo []*p2p.PeerInfo
	// Call function will process and return the rpc response into array.
	err = client.Call(&peerInfo, "admin_peers")
	if err != nil {
		log.Fatal(err)
	}
	//Print desired information 
	for _, peer := range peerInfo {
		fmt.Println("ID:", peer.ID)
		fmt.Println("Name:", peer.Name)
		// Access other properties as needed
	}
}

