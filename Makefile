.PHONY: check

SRC = $(shell find . -type f -name '*.go')

fmt:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -s -l logger.go example/main.go) || echo "Fix formatting issues with 'make fmt'"
	@golint ./...
	@go vet ./...
