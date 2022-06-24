# GO AMI v2 protocol implementation
#
.PHONY: test cov bench clean
all: test
	golint

test:
	go test -race -cover

cov:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out

codecov:
	go test -race -coverprofile=coverage.txt -covermode=atomic
	bash <(curl -s https://codecov.io/bash) -t f0414e7c-240f-492f-a78c-354bf33321d9

vet:
	@go vet -c=2

bench:
	@go test -bench=.

readme:
	mdr README.md

clean:
	rm -f coverage.out
	go clean
