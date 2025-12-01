.DEFAULT_GOAL := test-all

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

quality: ## Runs code quality checks
	@echo "=> Running code quality checks...."
	make lint
	make vet
	make fmt

lint: ## go linting
	@echo "=> Linting files...."
	golangci-lint run

vet: ## go vet
	@echo "=> Vetting files...."
	go vet ./...

fmt: ## go fmt
	@echo "=> Formatting files...."
	go fmt ./...

test-all: ## runs tests and generates coverage report
	@echo "=> Testing all...."
	go test -v -timeout 10m ./... -coverprofile=coverage.out -json > report.json

coverage: ## displays coverage output
	@echo "=> Opening HTML coverage report in browser..."
	make test-all
	go tool cover -html=coverage.out
	
clean: ## cleans generated files
	@echo "=> Removing binary and coverage files..."
	go clean
	rm coverage*.out

help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)