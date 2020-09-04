# GO AMI v2 protocol implementation
#
.PHONY: test cov bench clean

test:
	go test -race -cover

cov:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out

vet:
	@go vet -c=2

bench:
	@go test -bench=.

readme:
	mdr README.md

clean:
	rm -f coverage.out
	go clean
