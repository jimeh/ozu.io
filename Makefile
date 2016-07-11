DEV_DEPS = github.com/kardianos/govendor \
	github.com/vektra/mockery/.../ \
	github.com/mailru/easyjson/...

test: dev-deps
	@govendor test +local +program

generate: dev-deps
	@govendor generate +local +program

build:
	mkdir -p bin && go build -o bin/ozuio

run: build
	./bin/ozuio

fetch-vendor: dev-deps
	@govendor fetch +external +missing

install-vendor: dev-deps
	@govendor install +vendor

dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get $(DEP);)

update-dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get -u $(DEP);)

.PHONY: test generate build run install-vendor dev-deps update-dev-deps
