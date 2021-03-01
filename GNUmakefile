# PROVIDER_DIR is used instead of PWD since docker volume commands can be dangerous to run in $HOME.
# This ensures docker volumes are mounted from within provider directory instead.
PROVIDER_DIR := $(abspath $(lastword $(dir $(MAKEFILE_LIST))))
TEST         := "$(PROVIDER_DIR)/dynatrace"
GOFMT_FILES  := $$(find $(PROVIDER_DIR) -name '*.go' |grep -v vendor)
WEBSITE_REPO := github.com/hashicorp/terraform-website
PKG_NAME     := dynatrace
OS_ARCH      := $(shell go env GOOS)_$(shell go env GOARCH)

ifneq ($(PWD),$(PROVIDER_DIR))
$(error "Makefile must be run from the provider directory")
endif

default: build

all: build depscheck fmtcheck test testacc test-compile tests-lint tests-lint-fix tools vet website-lint website-lint-fix

build: fmtcheck
	go install

depscheck:
	@echo "==> Checking source code with 'git diff'..."
	@git diff --check || exit 1
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@echo "==> Checking source code with go mod vendor..."
	@go mod vendor
	@git diff --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; exit 1)

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@./scripts/gofmtcheck.sh

test: fmtcheck
	go test $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test -v $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	mkdir -p ~/.terraform.d/plugins/dynatrace.com/com/dynatrace/1.0.0/$(OS_ARCH)
	go clean -cache
	go build -o ~/.terraform.d/plugins/dynatrace.com/com/dynatrace/1.0.0/$(OS_ARCH)/terraform-provider-dynatrace_1.0.0_$(OS_ARCH)
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

tests-lint: tools
	@echo "==> Checking acceptance test terraform blocks code with terrafmt..."
	@terrafmt diff -f ./dynatrace --check --pattern '*_test.go' --quiet || (echo; \
		echo "Unexpected differences in acceptance test HCL formatting."; \
		echo "To see the full differences, run: terrafmt diff ./dynatrace --pattern '*_test.go'"; \
		echo "To automatically fix the formatting, run 'make tests-lint-fix' and commit the changes."; \
		exit 1)

tests-lint-fix: tools
	@echo "==> Fixing acceptance test terraform blocks code with terrafmt..."
	@find ./dynatrace -name "*_test.go" -exec sed -i ':a;N;$$!ba;s/fmt.Sprintf(`\n/fmt.Sprintf(`/g' '{}' \; # remove newlines for terrafmt
	@terrafmt fmt -f ./dynatrace --pattern '*_test.go'

tools:
	go install github.com/bflad/tfproviderdocs
	go install github.com/client9/misspell/cmd/misspell
	go install github.com/katbyte/terrafmt
	go mod tidy
	go mod vendor