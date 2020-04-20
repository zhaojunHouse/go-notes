# goimports
GOFMT_PRIVATE="gitlab.innotechx.com/shm"
GOFMT_LOCAL="gitlab.innotechx.com/shm/membership.shmiao.net"
GOFMT_FILES=$(shell find . -name '*.go' | grep -v \.pb\.go$ | xargs)
GOTEST_FILES=$(shell find . -name '*_test.go' | xargs)

.PHONY: all
all: fmt lint build vet

.PHONY: fmt
fmt:
	@goimports -l -w -private "${GOFMT_PRIVATE}" -local "${GOFMT_LOCAL}" ${GOFMT_FILES}

.PHONY: lint
lint:
	@golint ./...

.PHONY: build
build:
	@go build ./...

.PHONY: vet
vet:
	@go vet ./...

.PHONY: test
test:
	@go test ${GOTEST_FILES}