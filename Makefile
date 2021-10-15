build-proto:
	protoc --go_out=. --go-grpc_out=. proto/communicate.proto
build:
	go build -o bin/thanos .