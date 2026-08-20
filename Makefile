# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=artifactory-exporter
CMDPATH := .

all: test build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_NAME) -v $(CMDPATH)

test:
	$(GOTEST) -v ./...

## testcov: Run all tests and generate an HTML coverage report.
testcov:
	$(GOTEST) ./... -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

## vet: Run go vet against the code.
vet:
	$(GOCMD) vet ./...

## lint: Run golangci-lint against the code.
lint:
	golangci-lint run -v --timeout=15m ./...

## dep: Verify and tidy module dependencies.
dep:
	$(GOCMD) mod tidy
	$(GOCMD) mod verify

clean:
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_NAME)
	rm -f coverage.out coverage.html

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_NAME) -v $(CMDPATH)

build-docker-linux:
	docker build . -t artifactory-exporter:latest
