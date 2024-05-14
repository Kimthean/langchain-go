build: 
	@go build -o bin/langchain-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/langchain-go