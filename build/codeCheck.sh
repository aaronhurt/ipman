#!/usr/bin/env bash

## ensure golangci-lint - https://golangci-lint.run
golangci="$(go env GOPATH)/bin/golangci-lint"
if ! ${golangci} version &> /dev/null; then
  echo "Installing golangci-lint binary to ${golangci}"
  curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin
  echo "done"
fi

## run golangci-lint - all config in .golangci.yml
lintRun=$(${golangci} run --show-stats ./...)
lintRet=$?
echo "golangci-lint ${lintRun}"
exit $lintRet
