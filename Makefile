PROJECTNAME=$(shell basename "$(PWD)")

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=publicapi
BINARY_LINUX=$(BINARY_NAME)_linux

all: help

## go-clean  Delete last bin
go-clean: 
	@echo "  >  Deleting last bin"
	rm -rf ./bin/*

## go-run  Run binary
go-run:
	@echo "  >  Running binary..."
	./bin/$(BINARY_NAME)

## build  Create binary 
build:
	@echo "  >  Building binary..."
	$(GOBUILD) -o ./bin/$(BINARY_NAME) cmd/*.go

## build-linux  Build binary (Linux OS)
build-linux:
	@echo "  >  Building binary for Linux OS..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_LINUX) cmd/*.go

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: help
