# PeerConnect: Golang-based Peer Information Retrieval for GETH
This program used go-Ethereum library functions to interact with GETH client, and will print out a list of connected peers. 


Run Geth execution client
with jwt
```
geth --http --http.api eth,net,engine,admin --authrpc.jwtsecret /Your/Path/to/jwt.hex
```
without jwt
```
geth --http --http.api eth,net,engine,admin
```
GETH execution client HTTP port: 8545(default)

run the golang program.
