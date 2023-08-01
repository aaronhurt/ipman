#!/usr/bin/env bash

## ensure we have staticcheck
## https://github.com/golang/lint
if ! staticcheck=$(type -p "${GOPATH}/bin/staticcheck"); then
	echo -n "Installing staticcheck ... "
	go install honnef.co/go/tools/cmd/staticcheck@latest
	echo "done"
	staticcheck=$(type -p "${GOPATH}/bin/staticcheck")
fi

## ensure we have the misspell tool
## https://github.com/client9/misspell
if ! misspell=$(type -p "${GOPATH}/bin/misspell"); then
	echo -n "Installing misspell ... "
	go install github.com/client9/misspell/cmd/misspell@latest
	echo "done"
	misspell=$(type -p "${GOPATH}/bin/misspell")
fi

## ensure we have the gocyclo tool
## https://github.com/fzipp/gocyclo
if ! gocyclo=$(type -p "${GOPATH}/bin/gocyclo"); then
	echo -n "Installing gocyclo ... "
	go install go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	echo "done"
	gocyclo=$(type -p "${GOPATH}/bin/gocyclo")
fi

## check formatting ignoring git and vendor
fmtTest=$(find . -name '*.go' -not -path './.git/*' -not -path './vendor/*' | xargs gofmt -l -s 2>&1)
if [ ! -z "$fmtTest" ]; then
	echo "gofmt         failed"
	echo "$fmtTest"
	exit 1
else
	echo "gofmt         succeeded"
fi

## run go vet ignoring vendor and the silly "Error" bug/feature
## https://github.com/golang/go/issues/6407
vetTest=$(go vet ./... 2>&1 | egrep -v '^vendor/|\s+vendor/|/vendor/|^exit\ status|\ possible\ formatting\ directive\ in\ Error\ call')
if [ ! -z "$vetTest" ]; then
	echo "go vet        failed"
	echo "$vetTest"
	exit 1
else
	echo "go vet        succeeded"
fi

## run staticcheck ignoring vendor
staticTest=$(${staticcheck} ./... 2>&1 | egrep -v '^vendor/|\s+vendor/|/vendor/')
if [ ! -z "$statitTest" ]; then
	echo "staticcheck   failed"
	echo "$staticTest"
	exit 1
else
	echo "staticcheck   succeeded"
fi

## check misspell ignoring git, vendor and 3rdparty
spellTest=$(find . -name '*' -not -path './.git/*' -not -path './vendor/*' -not -path './3rdparty/*' | xargs ${misspell} 2>&1 | echo)
if [ ! -z "$spellTest" ]; then
	echo "misspell      failed"
	echo "$spellTest"
	exit 1
else
	echo "misspell      succeeded"
fi

## check gocyclo ignoring git and vendor
cycloTest=$(find . -name '*.go' -not -path './.git/*' -not -path './vendor/*' | xargs ${gocyclo} -over 15 2>&1 | echo)
if [ ! -z "$cycloTest" ]; then
	echo "gocyclo       failed"
	echo "$cycloTest"
	exit 1
else
	echo "gocyclo       succeeded"
fi
