build:
	@go build -o bin/librarie

run: build
	@./bin/librarie

test:
	@go test -v ./...
