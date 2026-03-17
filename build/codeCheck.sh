#!/usr/bin/env bash
set -e -o pipefail

## ensure golangci-lint - https://golangci-lint.run
golangci="$(go env GOPATH)/bin/golangci-lint"
if ! ${golangci} version &> /dev/null; then
  echo "Installing golangci-lint binary to ${golangci}"
  curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin
fi

## run golangci-lint - all config in .golangci.yml
echo "Running golangci-lint ..."
${golangci} run --show-stats ./...

## ensure gvulncheck
govulncheck="$(go env GOPATH)/bin/govulncheck"
if ! ${govulncheck} -version &> /dev/null; then
  echo "Installing govulncheck binary to ${govulncheck}"
  go install golang.org/x/vuln/cmd/govulncheck@latest
fi

## run govulncheck
echo "Running govulncheck ..."
${govulncheck} -show verbose ./...
