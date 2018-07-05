.DEFAULT_GOAL = all

version  := $(shell git rev-parse --short HEAD)

name     := krex
package  := github.com/kris-nova/$(name)
packages := $(shell go list ./... | grep -v /vendor/)

.PHONY: all
all:: dependencies
all:: build

.PHONY: dependencies
dependencies::
	dep ensure

.PHONY: build
build::
	go build -o krex -ldflags "-X ${package}/cmd.version=${version}" .

.PHONY: test
test::
	go test -v $(packages)

.PHONY: bench
bench::
	go test -bench=. -v $(packages)

.PHONY: lint
lint::
	go vet -v $(packages)

.PHONY: check
check:: lint test

.PHONY: clean
clean::
	rm -f krex
