.PHONY: help build build-local up down logs ps test run
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build image for deploy
	docker build -t kijimad/gorun:${DOCKER_TAG} \
	--target release ./

build-local: ## Build image for local development
	docker-compose build --no-cache

up: ## Do docker compose up
	docker-compose up -d

down: ## Do docker compose down
	docker-compose down

logs: ## Tail docker compose logs
	docker-compose logs -f

ps: ## Check container status
	docker-compose ps

lint: ## Run lint
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v

run: ## run
	go run .

test: ## Run test
	go test -race -shuffle=on -v ./...

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
