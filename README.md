1. Client connect to websocket, pass jwt on query params
2. jwt is going to be sent auth service via GRPC
3. Client is verified and connect to websocket server then cache in redis/keydb
