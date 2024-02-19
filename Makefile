# Makefile for the ezd-api-go project
# Located in the project root

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ezd-api-go
BINARY_UNIX=$(BINARY_NAME)_unix

# Paths
SRC_DIR=./src
BINARY_PATH=$(SRC_DIR)/$(BINARY_NAME)

# Targets
all: test build
build: 
	$(GOBUILD) -o $(BINARY_PATH) -v $(SRC_DIR)
test: 
	cd $(SRC_DIR) && $(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_PATH)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_PATH) -v $(SRC_DIR)
	$(BINARY_PATH)
deps:
	cd $(SRC_DIR) && $(GOGET) -v ./...
watch:
	reflex -r '\.go$$' -s -- sh -c 'make run'
