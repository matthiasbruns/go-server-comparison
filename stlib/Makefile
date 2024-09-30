.PHONY: all
all: gen run

.PHONY: gen
gen: ## Generate code
	go run github.com/swaggo/swag/cmd/swag init
	go generate ./...
	go mod tidy

.PHONY: run
run: ## Run the application
	go run main.go
