#!/usr/bin/env bash

## ensure golangci-lint - https://golangci-lint.run
golangci="$(go env GOPATH)/bin/golangci-lint"
if ! ${golangci} version &> /dev/null; then
  echo "Installing golangci-lint binary to ${golangci}"
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
  sh -s -- -b "$(go env GOPATH)/bin" v1.57.1
  echo "done"
fi

## run golangci-lint - all config in .golangci.yml
lintRun=$(${golangci} run --show-stats ./...)
lintRet=$?
echo "golangci-lint ${lintRun}"
exit $lintRet
