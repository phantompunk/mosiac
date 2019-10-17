deps:
	go get -u ./...
clean:
	rm -rf main
build:
	GOOS=linux GOARCH=amd64 go build main.go
local run:
	go run main.go