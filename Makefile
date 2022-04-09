.PHONY: build run
build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/$(SERVICE) cmd/$(SERVICE)/*.go
run: 
	./bin/$(SERVICE)