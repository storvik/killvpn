SHELL := /bin/bash
APPNAME = killvpn
PACKAGE = github.com/storvik/killvpn/${APPNAME}

COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-w \
    -X main.BuildTime=${BUILD_TIME} \
    -X main.CommitHash=${COMMIT_HASH}"

.PHONY: all clean vendor fmt lint help
.DEFAULT_GOAL := help

all: build

clean:
	go clean
	rm -rf build/

createbuild:
	mkdir -p build

vendor: ## Install govendor, go-bindata and sync dependencies
	go get github.com/kardianos/govendor
	govendor sync ${PACKAGE}

build: createbuild vendor ## Build no.vaegamat binary
	go build -o build/${APPNAME} ${LDFLAGS}

install: createbuild vendor ## Install no.vaegamat binary
	go install ${LDFLAGS} ${PACKAGE}

release: createbuild vendor ## Build no.vaegamat release binary for linux server
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/${APPNAME}_linux_amd64

fmt: ## Run gofmt linter
	@for d in `govendor list -no-status +local | sed 's/github.com.storvik.killvpn/./'` ; do \
		if [ "`gofmt -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ improperly formatted go files" && echo && exit 1; \
		fi \
	done

lint: ## Run golint linter
	@for d in `govendor list -no-status +local | sed 's/github.com.storvik.killvpn/./'` ; do \
		if [ "`golint $$d | tee /dev/stderr`" ]; then \
			echo "^ golint errors!" && echo && exit 1; \
		fi \
	done

check-vendor: ## Verify that vendored packages match git HEAD
	@git diff-index --quiet HEAD vendor/ || (echo "check-vendor target failed: vendored packages out of sync" && echo && git diff vendor/ && exit 1)

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
