SUDO ?=
RELEASE_VERSION = $(shell grep RELEASE_VERSION= build/build.sh | grep -oE '[0-9]+?\.[0-9]+?')

ifeq ($(SUDO),true)
	sudo = sudo
endif

.PHONY: build test release check clean distclean docker docker-release

export GO111MODULE = on

build:
	@build/build.sh

install:
	@build/build.sh -i

test:
	@go test -v ./...

release:
	@build/build.sh -r

check:
	@build/codeCheck.sh

clean:
	@build/build.sh -d

distclean:
	@build/build.sh -dc
