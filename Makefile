install:
	go get .
	go install github.com/hexira/go-ignore-cov@v0.3.0

build:
	go build -v ./...

test:
	go test -coverprofile coverage.out -covermode count -coverpkg=./... -v ./...

coverage:
	go-ignore-cov --file coverage.out
	go tool cover -func=coverage.out

lint:
	golangci-lint run
