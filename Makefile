VERSION = $(shell godzil show-version)
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X github.com/Songmu/gokoku.revision=$(CURRENT_REVISION)"
ifdef update
  u=-u
endif

export GO111MODULE=on

.PHONY: deps
deps:
	go get ${u} -d
	go mod tidy

.PHONY: devel-deps
devel-deps: deps
	go install golang.org/x/lint/golint@latest
	go install github.com/mattn/goveralls@latest
	go install github.com/Songmu/godzil/cmd/godzil@latest

.PHONY: test
test: deps
	go test

.PHONY: lint
lint: devel-deps
	golint -set_exit_status

.PHONY: cover
cover: devel-deps
	goveralls

.PHONY: release
release: devel-deps
	godzil release
