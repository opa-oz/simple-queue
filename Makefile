.PHONY: help

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(patsubst %/,%,$(dir $(mkfile_path)))

include .env
export $(shell sed 's/=.*//' .env)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

scheduler: ## Run scheduler locally
	go run cmd/scheduler/main.go

.PHONY: scheduler

redis: redis-stop ## Run Redis in Docker for local development
	docker run --name queue-redis -p 6379:6379 -d redis

.PHONY: redis

redis-stop: ## Stop local Redis in Docker
	docker stop queue-redis
	docker rm queue-redis

.PHONY: redis-stop
