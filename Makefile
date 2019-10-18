deps:
	go get -u ./...
clean:
	rm -rf main main.zip
build:
	GOOS=linux GOARCH=amd64 go build main.go
zip:
	zip main.zip main vendor/
local run:
	go run main.go