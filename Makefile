build:
	protoc --go_out=relative,plugins=grpc:./ ./client/test.proto
	protoc --go_out=relative,plugins=grpc:./ ./server/test.proto

server:
	go run ./server/*

client:
	go run ./client/*

.PHONY: server client
