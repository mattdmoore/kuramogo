install:
	go get .
	go install github.com/quantumcycle/go-ignore-cov@v0.4.0

build:
	go build -v ./...

test:
	go test -coverprofile coverage.out -covermode count -coverpkg=./... -v ./...

coverage:
	go-ignore-cov --file coverage.out
	go tool cover -func=coverage.out

lint:
	golangci-lint run
