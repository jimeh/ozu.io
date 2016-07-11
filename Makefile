DIRS = $(shell go list ./... | grep -v /vendor/)

DEV_DEPS = github.com/kardianos/govendor \
	github.com/vektra/mockery/.../ \
	github.com/mailru/easyjson/...

test:
	@go test $(DIRS)

generate: dev-deps
	@go generate $(DIRS)

build:
	mkdir -p bin && go build -o bin/ozuio

run: build
	./bin/ozuio

install-vendor:
	go install ./vendor/...

dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get $(DEP);)

update-dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get -u $(DEP);)

.PHONY: test build generate run install-vendor dev-deps
