GOPATH := $(shell go env GOPATH)

install:
	go install

uninstall-linux:
	rm -i ${GOPATH}/bin/ghost
