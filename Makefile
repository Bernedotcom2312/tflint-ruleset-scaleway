PLUGIN_NAME := tflint-ruleset-scaleway
PLUGIN_DIR  := $(HOME)/.tflint.d/plugins

.PHONY: default test e2e build install fmt vet catalog

default: build

test:
	go test $$(go list ./... | grep -v /integration)

e2e:
	cd integration && go test -count=1 ./...

build:
	go build

install: build
	mkdir -p $(PLUGIN_DIR)
	mv ./$(PLUGIN_NAME) $(PLUGIN_DIR)

fmt:
	gofmt -s -w .

vet:
	go vet ./...


# Refresh rules/models/catalog.go from the Scaleway API.
# Requires SCW_SECRET_KEY to be set.
catalog:
	go run ./tools/catalog -output rules/models/catalog.go
