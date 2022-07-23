

all:
	GOOS=linux  GOARCH=amd64 go build -o dobato-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o dobato-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o dobato-darwin-arm64
