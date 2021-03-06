DEV_DEPS = github.com/kardianos/govendor \
	github.com/vektra/mockery/.../ \
	github.com/mailru/easyjson/... \
	github.com/jteeuwen/go-bindata/... \
	github.com/mitchellh/gox

BINNAME = ozuio
BINARY = bin/${BINNAME}
DOCKERBIN = bin/ozuio_linux_amd64
DOCKERREPO = jimeh/ozu.io
BINDIR = $(shell dirname ${BINARY})
SOURCES = $(shell find . -name '*.go' -o -name 'VERSION')
VERSION = $(shell cat VERSION)
OSARCH = "darwin/386 darwin/amd64 linux/386 linux/amd64 linux/arm"
RELEASEDIR = releases

.DEFAULT_GOAL: test

.PHONY: test
test: dev-deps
	@govendor test +local +program

$(BINARY): $(SOURCES)
	go build -o ${BINARY} -ldflags "-X main.Version=${VERSION}"

.PHONY: build
build: $(BINARY)

.PHONY: clean
clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi; \
	if [ -f ${DOCKERBIN} ]; then rm ${DOCKERBIN}; fi; \
	if [ -d ${BINDIR} ]; then rmdir ${BINDIR}; fi

.PHONY: run
run: $(BINARY)
	$(BINARY)

.PHONY: install
install: dev-deps
	@govendor install +local +program

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

.PHONY: release-build
release-build:
	gox -output "${RELEASEDIR}/${BINNAME}_${VERSION}_{{.OS}}_{{.Arch}}" \
		-osarch=${OSARCH} \
		-ldflags "-X main.Version=${VERSION}"

.SILENT: release
.PHONY: release
release: release-build
	$(eval BINS := $(shell cd ${RELEASEDIR} && find . \
		-name "${BINNAME}_${VERSION}_*" -not -name "*.tar.gz"))
	cd $(RELEASEDIR); \
	$(foreach BIN,$(BINS),tar -cvzf $(BIN).tar.gz $(BIN) && rm $(BIN);)

$(DOCKERBIN): $(SOURCES)
	CGO_ENABLED=0 GOOS=linux ARCH=amd64 go build \
		-a -o ${DOCKERBIN} -ldflags "-X main.Version=${VERSION}"

.PHONY: build-docker
build-docker: $(DOCKERBIN)
	docker build -t "${DOCKERREPO}:latest" . \
	&& docker tag "${DOCKERREPO}:latest" "${DOCKERREPO}:${VERSION}"
