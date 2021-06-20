BINDIR      := $(CURDIR)/bin
BINNAME     ?= manifest-gen

GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif
GOX           = $(GOBIN)/gox
GOIMPORTS     = $(GOBIN)/goimports
ARCH          = $(shell uname -p)

# go option
PKG        := ./...
TAGS       :=
TESTS      := .
TESTFLAGS  :=
LDFLAGS    := -w -s
GOFLAGS    :=


.PHONY: all
all: lint build

# ------------------------------------------------------------------------------
#  lint

.PHONY: lint
lint:
	golangci-lint run --timeout 2m0s

# ------------------------------------------------------------------------------
#  build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)'/$(BINNAME) ./cmd/manifest-gen


# ------------------------------------------------------------------------------
#  install

.PHONY: install
install: build
	@install "$(BINDIR)/$(BINNAME)" "$(GOBIN)/$(BINNAME)"
