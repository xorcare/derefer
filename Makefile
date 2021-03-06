# Makefile for golang library project

# Copyright © 2019 Vasiliy Vasilyuk. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Based on https://git.io/fjkGc

# NAMESPACE the full path to the main package is use in the
# imports tool to format imports correctly.
NAMESPACE = github.com/xorcare/pointer

# COVER_FILE the name of the file recommended in the standard
# documentation go test -cover and used codecov.io
# to check code coverage.
COVER_FILE ?= coverage.out

# AT addition to commands to hide unnecessary command output.
AT ?= @

build: ## Build the project binary
	$(AT)go build ./...

check: static test build ## Check project with static checks and unit tests

ci: check ## Target for integration with ci pipeline

help: ## Print this help
	$(AT)grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

imports: tools ## Check and fix import section by import rules
	$(AT)test -z $$(for d in $$(go list -f {{.Dir}} ./...); \
	do goimports -e -l -local $(NAMESPACE) -w $$d/*.go; done)

lint: tools ## Check the project with lint
	$(AT)golint -set_exit_status ./...

static: imports vet lint ## Run static checks (lint, imports, vet, etc.) all over the project

test: ## Run unit tests
	$(AT)go test ./... -count=1 -race
	$(AT)go test ./... -count=1 -coverprofile=$(COVER_FILE) -covermode=atomic $d
	$(AT)go tool cover -func=$(COVER_FILE) | grep ^total

CDT ?= $(AT)cd internal/tools &&
tools: ## Install all needed tools, e.g., for static checks
	$(CDT) \
	go install \
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports

toolsup: ## Update all needed tools, e.g., for static checks
	$(CDT) \
	go mod tidy && \
	go get \
	golang.org/x/lint/golint$(AT)latest \
	golang.org/x/tools/cmd/goimports$(AT)latest && \
	go mod download && \
	go mod verify
	$(AT)make tools

vet: ## Check the project with vet
	$(AT)go vet ./...

.PHONY: build check ci help imports lint static test tools toolsup vet
.DEFAULT_GOAL := build
