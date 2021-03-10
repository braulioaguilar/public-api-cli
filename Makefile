PROJECTNAME=$(shell basename "$(PWD)")

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=publicapi
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_WINDOWS=$(BINARY_NAME)_windows.exe
BINARY_MAC=$(BINARY_NAME)_mac

## Delete last bin
clean: 
	@echo "  >  Deleting last bin"
	rm -rf ./bin/*

## Build binaries
build:
	@echo "  >  Building binary for Linux OS..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_LINUX) cmd/*.go

	@echo "  >  Building binary for Mac OS..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_MAC) cmd/*.go

	@echo "  >  Building binary for Windows OS..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_WINDOWS) cmd/*go