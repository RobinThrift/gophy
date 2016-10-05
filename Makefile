# gophy make targets
PREFIX?=$(shell pwd)

.PHONY: clean test build

build:
	@echo "+ $@"
	@go build .

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v vendor | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | grep -v vendor | tee /dev/stderr

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v vendor)

test: fmt lint vet
	@echo "+ $@"
	@go test -v -tags $(shell go list ./... | grep -v vendor)

clean:
	@echo "+ $@"
	@rm -rf gophy

install:
	@echo "+ $@"
	@go install .

