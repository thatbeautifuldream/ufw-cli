# Binary name
BINARY_NAME=ufw-cli

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

# Build directory
BUILD_DIR=build

# Main entry point
MAIN_FILE=main.go

.PHONY: all build clean run install

all: clean build

build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

run:
	$(GORUN) $(MAIN_FILE)

install: build
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

# Help target
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  clean    - Clean build files"
	@echo "  run      - Run the application"
	@echo "  install  - Install the binary to /usr/local/bin"
	@echo "  all      - Clean and build"