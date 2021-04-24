BINARY       ?= hibp
BUILD_FLAGS  ?=
TEST_TIMEOUT ?=120s
CURRENT_DIR  = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
VERSION      ?= $(shell git describe --tags --always)
LDFLAGS      ?= -X github.com/radekg/hibp/config.Version=$(VERSION) -w -s

SWAGGER_VERSION := v0.26.1
SWAGGER := docker run -u $(shell id -u):$(shell id -g) --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) -e GOCACHE=/tmp/.cache --entrypoint swagger quay.io/goswagger/swagger:$(SWAGGER_VERSION)

.PHONY: lint
lint:
	golint ./...

generate-server-api:
	mkdir -p api
	$(SWAGGER) generate server -s server -a restapi \
			-t api \
			-f .swagger/api.swagger.yaml \
			--exclude-main \
			--default-scheme=http