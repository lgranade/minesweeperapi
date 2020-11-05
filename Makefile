MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))

-include $(MKFILE_DIR)/config_deploy.env
-include $(MKFILE_DIR)/config_run.env

DOCKER_REPO=$(AWS_ACCOUNT_NR).dkr.ecr.$(AWS_REGION).amazonaws.com

APP_NAME=minesweeperapi

ifeq ($(OVERRIDE_VERSION), )
 	GIT_TAG=$(shell git describe --exact-match --tags 2> /dev/null)
	ifeq ($(GIT_TAG), )
		GIT_BRANCH_RAW=$(shell git branch | grep \*)
		ifeq ($(findstring no branch, $(GIT_BRANCH_RAW)), no branch)
			VERSION=unknown
		else
			GIT_BRANCH_ESCAPED=$(shell git branch | grep \* | cut -d ' ' -f2 | sed 's|/|-|g')
			VERSION=$(GIT_BRANCH_ESCAPED)
		endif
	else
		VERSION=$(GIT_TAG)
	endif
else
	VERSION=$(OVERRIDE_VERSION)
endif

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: build-image ## Default build

# Local

generate-sql: ## Generate SQL
	@cd $(MKFILE_DIR)/sql; sqlc generate

build-local: ## Build binary locally
	@mkdir -p $(MKFILE_DIR)/bin
	cd $(MKFILE_DIR)/src; go build -o $(MKFILE_DIR)/bin/minesweeperapi

run-local: ## Run binary locally
	DB_HOST=$(DB_HOST) \
	DB_PORT=$(DB_PORT) \
	DB_USER=$(DB_USER) \
	DB_PASSWORD=$(DB_PASSWORD) \
	DB_NAME=$(DB_NAME) \
	DB_SSLMODE=$(DB_SSLMODE) \
	$(MKFILE_DIR)/bin/minesweeperapi

unit-test-local: ## Run unit tests on all packages locally
	@cd $(MKFILE_DIR)/src; go test -v --tags=unit ./...

# Docker

build-image: ## Build deployable image
	docker build \
		--target "runner" -t $(APP_NAME) .

start-container: ## Run deployable image
	docker run -d --rm \
		-p $(API_PORT):8080 \
		--env-file="$(MKFILE_DIR)/config_run.env" \
		--net=host \
		--name="$(APP_NAME)" $(APP_NAME)

logs-container: ## Follow container logs
	docker logs -f $(APP_NAME)

stop-container: ## Stop and remove a running container
	-docker stop $(APP_NAME)

release: build publish ## Make a release by building and publishing the `{version}` and `latest` tagged containers to ECR

publish: publish-latest publish-version ## Publish the `{version}` ans `latest` tagged containers to ECR

publish-latest: tag-latest ## Publish the `latest` taged container to ECR
	@echo 'publish latest to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(APP_NAME):latest

publish-version: tag-version ## Publish the `{version}` taged container to ECR
	@echo 'publish $(VERSION) to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

tag: tag-latest tag-version ## Generate container tags for the `{version}` ans `latest` tags

tag-latest: ## Generate container `latest` tag
	@echo 'create tag latest'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):latest

tag-version: ## Generate container `{version}` tag
	@echo 'create tag $(VERSION)'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

version: ## Output the current version
	@echo $(VERSION)