.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help
SHELL := /usr/bin/env bash
SHELL_OPTIONS ?= -euo pipefail

avalanche: ## Start the Avalanche node
	go run . updateConfig ava
	docker compose -f ./network/compose-avalanche.yml up -d

klaytn: ## Start the Klaytn node
	go run . updateConfig klay
	docker compose -f ./network/compose-klaytn.yml up -d

ethereum: ## Start the Ethereum node
	go run . updateConfig eth
	docker compose -f ./network/compose-ethereum.yml up -d

build-mac: ## Build the project on MAC
	GOOS=darwin GOARCH=amd64 go build -o antps .

build-linux: ## Build the project on Linux
	GOOS=linux GOARCH=amd64 go build -o antps .

ava-output: ## Generate Avalanche output
	cd result && python graph.py ava

eth-output: ## Generate Ethereum output
	cd result && python graph.py eth

klay-output: ## Generate Klaytn output
	cd result && python graph.py klay
