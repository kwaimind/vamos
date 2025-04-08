build:
	go build && mv vamos ~/go/bin/

run:
	go run main.go

test:
	go test ./...

lint:
	golangci-lint run