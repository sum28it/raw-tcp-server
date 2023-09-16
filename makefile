tidy:
	go mod tidy
server:
	go run ./server/server.go

client:
	go run ./client/client.go