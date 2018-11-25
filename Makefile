PREFIX?=$(shell pwd)
NAME := base-go
BUILDDIR := ${PREFIX}/cross
VERSION := $(shell cat VERSION.txt)
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif
GO := go
GO_VERSION := $(shell grep golang .tool-versions | awk '{print $$2}')
DOCKER := docker
DOCKERUSER := $(shell whoami)
SED := $(shell which gsed || which sed)
MD5 := $(shell which gmd5sum || which md5sum)
SHA256 := $(shell which gsha256sum || which sha256sum)
GOOSARCHES = linux/amd64 darwin/amd64
CTIMEVAR=-X $(PKG)/version.GITCOMMIT=$(GITCOMMIT) -X $(PKG)/version.VERSION=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"

all: clean fmt test build install ## Runs a clean, fmt, test, build, and install

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	$(RM) $(NAME)
	$(RM) -r $(BUILDDIR)
	-@$(DOCKER) rm -f $(NAME)-postgres

.PHONY: deps
deps: ## Installs all dependencies
	go get -d -v ./..
	go get -u golang.org/x/lint/golint
	go get -u github.com/shuLhan/go-bindata/...
	go get -u github.com/kisielk/errcheck

.PHONY: fmt
fmt: ## Verifies all files have been `gofmt`ed
	@echo "+ $@"
	@$(GO) fmt ./...

.PHONY: test
test: ## Runs all tests
	@echo "+ $@"
	@golint ./...
	@$(GO) vet ./...
	@errcheck -asserts -blank ./...
	@$(GO) test -cover -coverprofile=coverage.out -v -tags "$(BUILDTAGS) cgo" ./...

$(NAME): $(wildcard *.go) $(wildcard */*.go) VERSION.txt
	@echo "+ $@"
	@$(GO) build -tags "$(BUILDTAGS)" ${GO_LDFLAGS} -o $(NAME) .

.PHONY: embeds
embeds: ## Embeds static assets into Go source
	@$(GO) generate ./...

.PHONY: build
build: embeds $(NAME) ## Builds a dynamic executable or package

.PHONY: run
run: ## Run main
	@echo "+ $@"
	$(GO) run main.go

.PHONY: dev
dev: ## Watch source files and run tests and main on save
	@echo "+ $@"
	@$(DOCKER) run --name $(NAME)-postgres -p 5432:5432 -d postgres:10
	@ag -l | entr -scrd 'make fmt test run'

.PHONY: install
install: ## Installs the executable or package
	@echo "+ $@"
	$(GO) install -a -tags "$(BUILDTAGS)" ${GO_LDFLAGS} .

define buildrelease
GOOS=$(1) GOARCH=$(2) CGO_ENABLED=0 $(GO) build \
	 -o $(BUILDDIR)/$(NAME)-$(1)-$(2) \
	 -a -tags "$(BUILDTAGS) static_build netgo" \
	 -installsuffix netgo ${GO_LDFLAGS_STATIC} .;
$(MD5) $(BUILDDIR)/$(NAME)-$(1)-$(2) > $(BUILDDIR)/$(NAME)-$(1)-$(2).md5;
$(SHA256) $(BUILDDIR)/$(NAME)-$(1)-$(2) > $(BUILDDIR)/$(NAME)-$(1)-$(2).sha256;
endef

.PHONY: ci
ci: ## Runs test suite in Docker build
	docker build --pull -f ci.dockerfile --build-arg GO_VERSION=$(GO_VERSION) .

.PHONY: release
release: *.go VERSION.txt ## Builds the cross-compiled binaries, naming them in such a way for release (eg. binary-GOOS-GOARCH)
	@echo "+ $@"
	$(foreach GOOSARCH,$(GOOSARCHES), $(call buildrelease,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH))))
	@$(DOCKER) build --pull -t $(DOCKERUSER)/$(NAME):$(GITCOMMIT) --build-arg SERVICE_NAME=$(NAME) $(BUILDDIR)
	@$(DOCKER) tag $(DOCKERUSER)/$(NAME):$(GITCOMMIT) $(DOCKERUSER)/$(NAME):$(VERSION)

.PHONY: bump-version
BUMP := patch
bump-version: ## Bump the version in the version file. Set BUMP to [ patch | major | minor ]
	@$(GO) get -u github.com/jessfraz/junk/sembump # update sembump tool
	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))
	@echo "Bumping VERSION.txt from $(VERSION) to $(NEW_VERSION)"
	echo $(NEW_VERSION) > VERSION.txt
	@echo "Updating links to download binaries in README.md"
	@$(SED) -i s/$(VERSION)/$(NEW_VERSION)/g README.md
	git add VERSION.txt README.md
	git commit -vsam "Bump version to $(NEW_VERSION)"
	@echo "Run make tag to create and push the tag for new version $(NEW_VERSION)"

.PHONY: tag
tag: ## Create a new git tag to prepare to build a release
	git tag -sa $(VERSION) -m "$(VERSION)"
	@echo "Run git push origin $(VERSION) to push your new tag to GitHub and trigger a build."

