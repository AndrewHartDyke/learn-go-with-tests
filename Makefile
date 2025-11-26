.DEFAULT_GOAL := test-all

test-all:
	@echo "=> Testing all...."
	go test ./...

test-all-verbose:
	go test ./... -v

cover:
	@echo "=> Running tests and generating coverage profile (coverage.out)..."
	go test -coverprofile=coverage.out ./...

html-cover: cover
	@echo "=> Opening HTML coverage report in browser..."
	go tool cover -html=coverage.out

cover-remove:
	@echo "=> Removing coverage files..."
	@rm coverage.out