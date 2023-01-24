.PHONY: test
test:
	go test ./...

build-arm:
	mkdir -p bin
	env GOOS=linux GOARCH=arm go build -o ./bin/srtorder-arm .
