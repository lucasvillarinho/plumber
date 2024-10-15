.PHONY: test
test: ## Run unit test
	go test -v -coverprofile=rawcover.out -json $$(go list ./... ) 2>&1 | tee /tmp/gotest.log | gotestfmt -hide successful-tests,empty-packages
