# Run tests.
test:
	bash -c "set -m; bash '$(CURDIR)/scripts/test.sh'"

# Check code quality.
lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6 run
	npx prettier . --check

format-main:
	npx prettier . --write

format-zerolog:
	cd "$(CURDIR)/zerolog" && go mod tidy
	cd "$(CURDIR)/zerolog" && go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6 run --fix

format-golm:
	cd "$(CURDIR)/golm" && go mod tidy
	cd "$(CURDIR)/golm" && go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6 run --fix

format-sentry:
	cd "$(CURDIR)/sentry" && go mod tidy
	cd "$(CURDIR)/sentry" && go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6 run --fix

# Reformat code so it passes the code style lint checks.
format: format-main format-zerolog format-golm format-sentry
