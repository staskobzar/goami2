# GO AMI v2 protocol implementation
#
.PHONY: test cov bench clean parser
all: test

dev: test parser

test: parser
	go test -race -cover -timeout=3s

parser:
	re2go parse.re -o parse.go -i --no-generation-date

cov: parser
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out

codecov:
	go test -race -coverprofile=coverage.txt -covermode=atomic
	bash <(curl -s https://codecov.io/bash) -t f0414e7c-240f-492f-a78c-354bf33321d9

lint:
	golangci-lint run

bench:
	@go test -benchmem -bench=.

clean:
	rm -f coverage.out
	go clean
