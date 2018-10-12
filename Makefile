generate:
	@go generate ./...
.PHONY: generate

install:
	@go get -u -v github.com/golang/dep/cmd/dep
	@dep ensure

build: generate
	@echo "====> Build telnetor cli"
	@go build -o ./bin/mailer-send-cloud main.go
.PHONY: build

release:
	@echo "====> Build and release"
	@go get github.com/goreleaser/goreleaser
	@git reset --hard HEAD
	@goreleaser
.PHONY: release

test:
	@go test ./...
.PHONY: test

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic
.PHONY: test.cov
