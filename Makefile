# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help build run
.DEFAULT_GOAL := help
CLIENT_NAME = vm
GOOS = $$(go env GOOS)
GOARCH = $$(go env GOARCH)

build: ## builds the vm go binary
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$(CLIENT_NAME)-$(GOOS)-$(GOARCH) cmd/vm/main.go

run: ## run vm backend
	@bin/$(CLIENT_NAME)-$(GOOS)-$(GOARCH)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
