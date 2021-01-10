.PHONY: tests

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

##
### Code validation
tests:
	cd clientrest && go test -v -race ./... && \
	cd .. && \
	cd basket && go test -v -race ./...