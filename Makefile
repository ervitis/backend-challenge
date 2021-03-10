.PHONY: tests dependencies help

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

tests: ## Test the client and backend
	cd clientrest && go test -v -race ./... && \
	cd .. && \
	cd basket && go test -v -race ./...

dependencies: ## Download dependencies
	go get -v -u ./...

client: ## Run client
	go run clientrest/cmd/main.go

server: ## Run server
	go run basket/cmd/main.go

compile: ## Compile client and server
	cd basket/cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o basketserver ./... && mv basketserver ../ && \
	cd ../.. && \
	cd clientrest/cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o basketclient ./... && mv basketclient ../