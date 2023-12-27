proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. message.proto 

run:
	go build -o service && ./service