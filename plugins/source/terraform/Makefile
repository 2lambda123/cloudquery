# Generate mocks for mock/unit testing 
.PHONY: generate-mocks
generate-mocks:
	go generate ./client/services/...

.PHONY: build
build:
	go build

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: gen-docs
gen-docs: build
	@command -v cloudquery >/dev/null 2>&1 || { \
		echo "Error: 'cloudquery' command not found. Please install it before running gen-docs."; \
		echo "You can install it by following the instructions at: https://www.cloudquery.io/docs/quickstart"; \
		exit 1; \
	}
	rm -rf ../../../website/tables/terraform
	cloudquery tables --format markdown --output-dir ../../../website/tables test/config.yml
	sed 's_(\(.*\))_(../../../../../website/tables/terraform/\1)_' ../../../website/tables/terraform/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/terraform/README.md
	mv ../../../website/tables/terraform/README.md ../../../website/pages/docs/plugins/sources/terraform/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/terraform/*.md
	rm -rf ../../../website/tables/terraform/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

# All gen targets
.PHONY: gen
gen: gen-docs
