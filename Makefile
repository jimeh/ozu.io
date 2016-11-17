DEV_DEPS = github.com/kardianos/govendor \
	github.com/vektra/mockery/.../ \
	github.com/mailru/easyjson/... \
	github.com/jteeuwen/go-bindata/... \
	github.com/mitchellh/gox

BINNAME = ozuio
BINARY = bin/${BINNAME}
BINDIR = $(shell dirname ${BINARY})
SOURCES = $(shell find . -name '*.go' -o -name 'VERSION')
VERSION = $(shell cat VERSION)
OSARCH = "linux/amd64 darwin/amd64"

.DEFAULT_GOAL: $(BINARY)
$(BINARY): $(SOURCES)
	go build -o ${BINARY} -ldflags "-X main.Version=${VERSION}"

.PHONY: build
build: $(BINARY)

.PHONY: clean
clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi; \
	if [ -d ${BINDIR} ]; then rmdir ${BINDIR}; fi

.PHONY: run
run: $(BINARY)
	$(BINARY)

.PHONY: install
install: dev-deps
	@govendor install +local +program

.PHONY: test
test: dev-deps
	@govendor test +local +program

.PHONY: generate
generate: dev-deps
	@govendor generate +local +program

.PHONY: vendor-sync
vendor-sync: dev-deps
	@govendor sync

.PHONY: vendor-fetch
vendor-fetch: dev-deps
	@govendor fetch +external +missing

.PHONY: vendor-install
vendor-install: dev-deps
	@govendor install +vendor

.PHONY: dev-deps
dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get $(DEP);)

.PHONY: update-dev-deps
update-dev-deps:
	@$(foreach DEP,$(DEV_DEPS),go get -u $(DEP);)

.PHONY: web-generate
web-generate:
	cd web && go generate

.PHONY: web-debug-bindata
web-debug-bindata:
	cd web && go-bindata -debug -pkg web static/... templates/...

.PHONY: package
package: dev-deps generate
	gox -output "pkg/${VERSION}/${BINNAME}_${VERSION}_{{.OS}}_{{.Arch}}" \
		-osarch=${OSARCH} \
		-ldflags "-X main.Version=${VERSION}" \
	&& gzip -9 pkg/${VERSION}/${BINNAME}_*

bin/ozuio_linux_amd64: $(SOURCES)
	CGO_ENABLED=0 GOOS=linux ARCH=amd64 \
		go build -a -o bin/ozuio_linux_amd64 \
		-ldflags "-X main.Version=${VERSION}"

.PHONY: build-docker
build-docker: bin/ozuio_linux_amd64
	docker build -t "jimeh/ozu.io:latest" . \
	&& docker tag "jimeh/ozu.io:latest" "jimeh/ozu.io:${VERSION}"
