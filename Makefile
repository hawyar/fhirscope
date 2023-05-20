,PHONY: build clean

name=fhirscope
bin = ./bin/$(name)
os = $(shell go env GOOS)
arch = $(shell go env GOARCH)

test:
	go test -v ./... -count=1

test-coverage:
	rm -rf ./coverage.out
	go test -v ./... -count=1 -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	GOOS=$(os) GOARCH=$(arch) go build -o $(bin) .
	chmod +x $(bin)

clean:
	rm -rf ./build ./data ./logs
	go clean
