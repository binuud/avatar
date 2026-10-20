# Maintainer Binu Udayakumar

# import config.
# You can change the default config with `make cnf="config_special.env" build`
# use build_staging.env for staging server (local too)

### WARNING ### PUBLIC-REPO ##########

# Optional environmental variables
cnf ?= .env

GITHUB_TOKEN=""
REPO=binuud
APP_NAME ?= avatar

BUILD_VER ?= v0.0.1
DOCKER_HUB_TAG ?= v0.0.1

GOLANG_IMAGE=golang:1.27.2

TAGGED_NAME = $(REPO)/$(APP_NAME)

MULTI_PLATFORM_DOCKER = --platform=linux/amd64,linux/arm64

## override any variable with .local.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))


# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: build-ui ## Build the avatar server binary, with embedded static html files
	go build cmd/avatar/main.go

build-docker: ## Build docker image
	docker build --no-cache  --build-arg BUILD_VER=$(BUILD_VER) $(MULTI_PLATFORM_DOCKER) -t $(REPO)/$(APP_NAME) -f deployment/Dockerfile .

build-ui: ## build ui image, and copy distribution file to ./ui folder
	cd ../avatar-dev-ui && make build
	rm -rf cmd/avatar/ui/*
	cp -R ../avatar-dev-ui/dist/avatar-dev-ui/browser/ cmd/avatar/ui/

deploy: ## Deploy all images to docker hub
	@echo "====================> Pushing image to dockerhub (PUBLIC) ."
## @docker login --username=$(DOCKER_HUB_USER) --password="$(DOCKER_HUB_PASSWD)"
	docker tag $(REPO)/${APP_NAME} $(REPO)/${APP_NAME}:$(DOCKER_HUB_TAG)
	docker push $(REPO)/${APP_NAME}:$(DOCKER_HUB_TAG)

deploy-latest: ## deploy current tag as latest
	docker tag $(REPO)/${APP_NAME} $(REPO)/${APP_NAME}:latest
	docker push $(REPO)/${APP_NAME}:latest	

vendor: ## go vendor and tidy
	echo "tidy + vendor"
	go mod tidy
	go mod vendor	

run: ## Run server, UI accessible in http://localhost:9080/
	go run cmd/avatar/main.go -v -grpc_port 9090 -http_port 9080

test: ## Run tests
	go test -v ./pkg/...	

release: ## for releasing the package
## comments for reference
# git tag -a v0.1.0 -m "Release COMMENT"
# git push origin [version]
# export GITHUB_TOKEN="YOUR_GH_TOKEN"
	goreleaser release --clean

release-check: ## run release build, but do not publish
	goreleaser check
	goreleaser release --snapshot --clean	