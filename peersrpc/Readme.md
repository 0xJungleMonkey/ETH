This program did not use the go-ethereum rpc functions (Dial to connect, call to get information), instead, it is finshed by http post call, and decode the response into result struct. 

### Summary
1. establish connection
2. the repsonse of API needs to be filled in the certain object in order to process it. 
