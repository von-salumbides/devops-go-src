.PHONY: build clean run
environment = $(DEPLOY_ENV)
function_name = $(FUNCTION_NAME)
build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/$(SERVICE) cmd/$(SERVICE)/*.go
clean:
	rm -rf ./bin
run: 
	go run $(SERVICE)