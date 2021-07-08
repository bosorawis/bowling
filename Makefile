.PHONY: help clean fmt test  build all


all:    ## clean, format, build and unit test
	make clean
	make gofmt
	make build
	make test


clean:  ## go clean
	rm -rf bin/

fmt:    ## format the go source files
	go fmt ./...

test:
	go test ./... -count=1

build:
	go build -o bin/bowling ./cmd
