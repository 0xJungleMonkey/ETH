package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	// Set the JWT secret
	jwtSecret := "/Users/xinqi/geth/consensus/jwt.hex" // Replace with the actual path to your JWT secret

	// Create a custom HTTP client with JWT authentication
	jwtClient := &JWTAuthClient{Secret: jwtSecret}
	httpClient := &http.Client{Transport: jwtClient}

	// Connect to the local Geth client with the custom HTTP client
	client, err := rpc.DialHTTPWithClient("http://localhost:8545", httpClient)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}

	// Monitor and output the list of connected peers
	connectedPeers := []*PeerInfo{}
	for {
		// Get the list of connected peers
		var peers []*PeerInfo
		err := client.Call(&peers, "admin_peers")
		if err != nil {
			log.Printf("Failed to retrieve the list of peers: %v\n", err)
		}

		// Check for changes in the list of connected peers
		removedPeers, addedPeers := diffPeerLists(connectedPeers, peers)

		// Output the changes in the list of connected peers
		if len(removedPeers) > 0 || len(addedPeers) > 0 {
			fmt.Println()
			// fmt.Println("Changes in Connected Peers:")
			// fmt.Println("Removed Peers:")
			// for _, peer := range removedPeers {
			// 	fmt.Printf("- ID: %s, Name: %s, Enode: %s\n", peer.ID, peer.Name, peer.Enode)
			// }
			// fmt.Println("Added Peers:")
			// for _, peer := range addedPeers {
			// 	fmt.Printf("+ ID: %s, Name: %s, Enode: %s\n", peer.ID, peer.Name, peer.Enode)
			// }

			fmt.Println("Newest List of Connected Peers:")
			for _, peer := range peers {
				fmt.Printf("* ID: %s, Name: %s, Enode: %s\n", peer.ID, peer.Name, peer.Enode)
			}
			fmt.Println()
		}

		// Update the stored list of connected peers
		connectedPeers = peers

		// Sleep for 10 seconds before retrieving the list again
		time.Sleep(10 * time.Second)
	}
}

// JWTAuthClient is a custom HTTP client that includes JWT authentication
type JWTAuthClient struct {
	Secret string
}

// RoundTrip sends an HTTP request with JWT authentication
func (c *JWTAuthClient) RoundTrip(req *http.Request) (*http.Response, error) {
	// Add the JWT authorization header
	token := generateJWT(c.Secret)
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute the request using the default HTTP client
	return http.DefaultTransport.RoundTrip(req)
}

// Generate a JWT token using the secret
func generateJWT(secret string) string {
	claims := jwt.MapClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(secret))
	return signedToken
}

// Compare two lists of connected peers and return the added and removed peers
func diffPeerLists(oldPeers, newPeers []*PeerInfo) (removedPeers, addedPeers []*PeerInfo) {
	oldSet := make(map[string]bool)
	for _, peer := range oldPeers {
		oldSet[peer.ID] = true
	}

	newSet := make(map[string]bool)
	for _, peer := range newPeers {
		newSet[peer.ID] = true
	}

	for _, peer := range oldPeers {
		if !newSet[peer.ID] {
			removedPeers = append(removedPeers, peer)
		}
	}

	for _, peer := range newPeers {
		if !oldSet[peer.ID] {
			addedPeers = append(addedPeers, peer)
		}
	}

	return removedPeers, addedPeers
}

