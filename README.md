1. Run Geth execution client

 with jwt
   ```bash
   geth --http --http.api eth,net,engine,admin --authrpc.jwtsecret /Your/Path/to/jwt.hex
   ```
 without jwt 
   ```bash
   geth --http --http.api eth,net,engine,admin 
   ```
GETH execution client HTTP port: 8545(default)

2. run the golang program.
```bash
go build main.go
./main
```

3. This program used go-Ethereum library functions to interact with GETH client, and will print out a list of connected peers. 

