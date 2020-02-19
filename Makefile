# .SILENT: ; # no need for @
PROJECT         =magda
PROJECT_DIR		=$(shell pwd)

OS              := $(shell go env GOOS)
ARCH            := $(shell go env GOARCH)

DST_DIR         =gen
PROTODIR        =protos
WORKDIR         :=$(PROJECT_DIR)/_workdir
BINDIR          :=$(WORKDIR)/bin
GOSWAGGER       :=$(BINDIR)/swagger

setup:
	mkdir -p $(WORKDIR)
	GOBIN=$(BINDIR) go install github.com/99designs/gqlgen
	cd www && yarn install

run:
	cd www && yarn serve