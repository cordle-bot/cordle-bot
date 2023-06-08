GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

DOCKER ?= docker

# clean & dev

clean:
	@rm -r build

dev:
	@bash scripts/dev.sh

## testing

test:
	$(GO) clean -testcache 
	$(GO) mod tidy
	$(GO) test -cover ./...

## deploy & build

build: 
	$(GO) build -o build/program/app cmd/cli/main.go 

tdeploy:
	$(DOCKER) build --tag cordle2 .
	$(DOCKER) run -it -rm cordle2

deploy:
	$(DOCKER) build --tag cordle2 .
	$(DOCKER) run -it cordle2

mysql:
	$(DOCKER) build --tag mysql -f deployment/mysql/Dockerfile .
	$(DOCKER) run -it -p 3306:3306 mysql

# fmt

fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: dev clean test build tdeploy deploy fmt mysql