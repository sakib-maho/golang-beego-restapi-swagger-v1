run:
	APP_ADDRESS=:8080 go run ./cmd/server

test:
	go test ./...
