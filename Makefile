
TESTFLAGS ?= -cover

test:
	@go test $(TESTFLAGS) ./...

bench:
	@go test -bench=. ./... -run notest

.PHONY: test bench