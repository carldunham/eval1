.PHONY: run test lint generate tidy check

run:
	go run main.go

test:
	go test ./... -v

lint:
	golangci-lint run

generate:
	go run github.com/99designs/gqlgen generate

tidy:
	go mod tidy

# Combined targets
check: lint test    # Run both lint and test
