# .SILENT: ; # no need for @
PROJECT         =magda
PROJECT_DIR		=$(shell pwd)

OS              := $(shell go env GOOS)
ARCH            := $(shell go env GOARCH)

DST_DIR         =gen
PROTODIR        =protos
WORKDIR         :=$(PROJECT_DIR)/_workdir
PROTODOT_URL    =https://protodot.seamia.net/binaries/darwin

setup:
	mkdir -p $(WORKDIR)
	curl "$(PROTODOT_URL)" > $(WORKDIR)/protodot
	chmod 755 $(WORKDIR)/protodot
	go get github.com/grpc-ecosystem/grpc-gateway

compile:
	mkdir -p $(DST_DIR)

	protoc \
		-I=$(PROTODIR) \
		-I=/usr/local/include \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:src/service \
		--doc_out=$(DST_DIR) \
		--doc_opt=markdown,docs.md \
		--grpc-gateway_out=logtostderr=true:src/service \
		${PROTODIR}/service.proto \
		${PROTODIR}/entry.proto \
		${PROTODIR}/file.proto \
		${PROTODIR}/entity.proto

	$(WORKDIR)/protodot -config protodot/config.json -src protos/service.proto -output magda

	rm -fR README.md
	cat README.raw.md > README.md
	cat gen/docs.md >> README.md