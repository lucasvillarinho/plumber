GO_FILES := $(shell find . -name '*.go')

GOFMT := gofumpt
GOIMPORTS := goimports

.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/ ./...
	@echo "Build completed successfully"


.PHONY: run 
run: ## Run the application
	go run main.go

.PHONY: lint
lint: build ## Run lint
	@echo "Running linter..."
	@golangci-lint run ./...
	@echo "Linter passed successfully"


.PHONY: test
test: ## Run unit test
	go test -v -coverprofile=rawcover.out -json $$(go list ./... ) 2>&1 | tee /tmp/gotest.log | gotestfmt -hide successful-tests,empty-packages
