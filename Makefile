build:
	go build -v ./...

test:
	go test ./... -cover -coverprofile coverage.out -race -mod=vendor -v

run:
	go run github.com/jeanbritz/go-android-bruteforce-pin.git/cmd

vendor:
	go mod tidy
	go mod vendor