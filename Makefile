export GO111MODULE=on
export GOPROXY=https://goproxy.io

default: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

lint: ## Apply go lint check
	@golangci-lint run --timeout 10m ./...
.PHONY: lint

build: ## Build CLI for this project
	@go mod tidy
	@go build -o cli/showstart cli/main.go
.PHONY: build
