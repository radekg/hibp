BINARY       ?= hibp
BUILD_FLAGS  ?=
TEST_TIMEOUT ?=120s
VERSION      ?= $(shell git describe --tags --always)

LDFLAGS      ?= -X github.com/radekg/hibp/config.Version=$(VERSION) -w -s

SWAGGER_VERSION := v0.26.1
SWAGGER := docker run -u $(shell id -u):$(shell id -g) --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) -e GOCACHE=/tmp/.cache --entrypoint swagger quay.io/goswagger/swagger:$(SWAGGER_VERSION)

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
		-o $(BINARY)-linux-amd64 \
		$(BUILD_FLAGS) \
		-ldflags "$(LDFLAGS)" ./main.go

.PNONY: docker-build
docker-build:
	docker build -t localhost/hibp:latest .

generate-server-api:
	mkdir -p api
	$(SWAGGER) generate server -s server -a restapi \
			-t api \
			-f .swagger/api.swagger.yaml \
			--exclude-main \
			--default-scheme=http