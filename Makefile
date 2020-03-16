# .SILENT: ; # no need for @
PROJECT         =magda
PROJECT_DIR		=$(shell pwd)

OS              := $(shell go env GOOS)
ARCH            := $(shell go env GOARCH)

ifeq ($(OS),darwin)
	MACOS_OR_LINUX    :=macos
else
	MACOS_OR_LINUX    :=linux
endif

WORKDIR         :=$(PROJECT_DIR)/_workdir
BINDIR          :=$(WORKDIR)/bin
WWW_DIR         :=$(PROJECT_DIR)/www
GCOULD_VERSION  :=283.0.0
GCLOUD_FILENAME :=google-cloud-sdk-$(GCOULD_VERSION)-$(OS)-x86_64.tar.gz
GCLOUD_URL      :=https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/$(GCLOUD_FILENAME)
GCLOUD_CMD      :=$(WORKDIR)/google-cloud-sdk/bin/gcloud

OVERMIND_VERSION  :=2.1.0
OVERMIND_FILENAME := overmind-v$(OVERMIND_VERSION)-$(MACOS_OR_LINUX)-amd64
OVERMIND_URL      := https://github.com/DarthSim/overmind/releases/download/v$(OVERMIND_VERSION)/$(OVERMIND_FILENAME).gz
OVERMIND_CMD      :=$(WORKDIR)/overmind


setup: workdir install_gcloud install_overmind
	mkdir -p $(WORKDIR)
	cd www && yarn install
	GOBIN=$(BINDIR) go install github.com/99designs/gqlgen
	GOBIN=$(BINDIR) go install golang.org/x/lint/golint

workdir:
	mkdir -p $(WORKDIR)

install_gcloud:
	curl $(GCLOUD_URL) > $(WORKDIR)/$(GCLOUD_FILENAME)
	cd $(WORKDIR) && tar -xzf $(GCLOUD_FILENAME)
	$(GCLOUD_CMD) components install --quiet beta

install_overmind:
	curl -L $(OVERMIND_URL) > $(WORKDIR)/$(OVERMIND_FILENAME).gz
	cd $(WORKDIR) && \
		gzip -d $(OVERMIND_FILENAME).gz && \
		mv $(OVERMIND_FILENAME) overmind && \
		chmod 755 overmind

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
	cd $(WWW_DIR) && FIRESTORE_EMULATOR_HOST=localhost:8972 yarn run test:firestore

gcloud_init:
	$(GCLOUD_CMD) init

overmind_dev:
	$(OVERMIND_CMD) start -f Procfile.test

overmind_test_run:
	$(OVERMIND_CMD) start -D -f Procfile.test

overmind_test_quit:
	$(OVERMIND_CMD) quit

test: overmind_test_run gotest golint jstestunit jstestfirebase overmind_test_quit

update:
	cd $(WWW_DIR) && yarn upgrade
	go get -u

run:
	$(BINDIR)/forego start

run_firestore:
	$(GCLOUD_CMD) beta emulators firestore start --quiet --host-port=localhost:8972 --rules=$(WWW_DIR)/firestore.rules