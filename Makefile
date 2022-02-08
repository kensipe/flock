
GOLANGCI_LINT_VER = "1.44.0"

.PHONY: build
build:
	go build ./...


# Run unit tests
.PHONY: test
test:
	go test ./...  -v -mod=readonly -coverprofile cover.out


.PHONY: lint
lint: ## Run golangci-lint
ifneq (${GOLANGCI_LINT_VER}, "$(shell ./bin/golangci-lint version --format short 2>&1)")
	@echo "golangci-lint missing or not version '${GOLANGCI_LINT_VER}', downloading..."
	curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/v${GOLANGCI_LINT_VER}/install.sh" | sh -s -- -b ./bin "v${GOLANGCI_LINT_VER}"
endif
	./bin/golangci-lint --timeout 3m run --build-tags integration