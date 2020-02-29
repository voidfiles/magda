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
FIREBASE_PATH   :=$(WORKDIR)/firebase

setup:
	mkdir -p $(WORKDIR)
	./bin/install_firebase.sh
	./bin/install_firebase_emulators.sh
	GOBIN=$(BINDIR) go install github.com/99designs/gqlgen
	cd www && yarn install

server:
	go run main.go

frontend:
	cd www && yarn serve

run:
	make server & make frontend