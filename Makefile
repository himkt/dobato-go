

all:
	GOOS=linux  GOARCH=amd64 go build -o bin/dobato-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/dobato-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o bin/dobato-darwin-arm64
