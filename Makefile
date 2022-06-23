GOPATH := $(shell go env GOPATH)

install:
	go install

uninstall:
	rm -i ${GOPATH}/bin/ghost
