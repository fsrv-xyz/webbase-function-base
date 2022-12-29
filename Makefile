.phony: build

build: export CGO_ENABLED = 0
build: export BUILD_FLAGS = -ldflags "-s -w"

REPLACEMENT_PATH ?= /project

build: replace main.go
	go build ${BUILD_FLAGS} -o /build/app /build/main.go

replace:
	go mod edit -replace github.com/fsrv-xyz/webbase-function-base/placeholder=$(REPLACEMENT_PATH)
	go mod tidy

