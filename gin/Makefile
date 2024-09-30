.PHONY: all
all: install gen run

.PHONY: install
install: ## Install dependencies
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: gen
gen: ## Generate code
	swag init
	go generate ./...
	go mod tidy

.PHONY: run
run: ## Run the application
	go run main.go
