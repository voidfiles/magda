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
WWW_DIR         :=$(PROJECT_DIR)/www
FIREBASE_PATH   :=$(WWW_DIR)/node_modules/.bin/firebase

setup:
	mkdir -p $(WORKDIR)
	cd www && yarn install
	./bin/install_firebase_emulators.sh
	GOBIN=$(BINDIR) go install github.com/99designs/gqlgen
	GOBIN=$(BINDIR) go install golang.org/x/lint/golint
	
generate:
	go run github.com/99designs/gqlgen generate

server:
	go run main.go

frontend:
	cd $(WWW_DIR) && yarn serve

golint:
	$(BINDIR)/golint main.go
	$(BINDIR)/golint pkg/...

gotest:
	go test ./...

jstestunit:
	cd $(WWW_DIR) && yarn run test:unit

jsteste2e:
	cd $(WWW_DIR) && yarn run test:e2e

jstestfirebase:
	cd $(WWW_DIR) && $(FIREBASE_PATH) emulators:exec --only=firestore "yarn run test:firestore"

test: gotest golint jstestunit jstestfirebase

run:
	make server & make frontend