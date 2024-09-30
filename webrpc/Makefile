.PHONY: all
all: gen run

.PHONY: gen
gen: ## Generate code
	go generate ./...
	go mod tidy

.PHONY: run
run: ## Run the application
	go run main.go

.PHONY: test
test: ## Test the application
	go test ./...
