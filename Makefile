# Run tests.
test:
	bash -c "set -m; bash '$(CURDIR)/scripts/test.sh'"

# Check code quality.
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

format-main:
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

format-zerolog:
	cd "$(CURDIR)/zerolog" && go mod tidy
	cd "$(CURDIR)/zerolog" && go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

# Reformat code so it passes the code style lint checks.
format: format-main format-zerolog
