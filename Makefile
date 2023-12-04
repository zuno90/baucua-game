test:
	go test
dev:
	nodemon --exec go run main.go --signal SIGTERM
client:
	go run grpc-client.go
proto:
	rm -f pb/*.go
	protoc --proto_path=../proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	../proto/auth.proto
