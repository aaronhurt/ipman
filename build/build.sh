#!/usr/bin/env bash
#
## package declarations
BUILD_NAME=$(basename "$(grep ^module go.mod|cut -f2 -d' ')")
RELEASE_VERSION="0.1"
RELEASE_BUILD=0

## default to build
GO_ACTION=build

## simple usage example
showUsage() {
  printf "Usage: %s [-d|-c|-r]
  -d    Remove binary and vendor directory
  -c    Clean distribution files
  -r    Build and package release binaries
  -i    Run 'go install' instead of 'go build'\n\n" "$0"
  exit 0
}

## install gox if needed
ensureGox() {
  if ! which gox &> /dev/null; then
    printf "Installing gox ... "
    go install github.com/mitchellh/gox@v1
  fi
}

## we require module support
export GO111MODULE=on

## exit toggle
should_exit=false

## read options
while getopts ":dcri" opt; do
  case $opt in
    d)
      printf "Removing binary and vendor directory ... "
      rm -rf "${BUILD_NAME}" vendor
      printf "done.\n"
      should_exit=true
      ;;
    c)
      printf "Cleaning dist directory ... "
      rm -rf ./dist/
      printf "done.\n"
      should_exit=true
      ;;
    r)
      ensureGox
      RELEASE_BUILD=1
      ;;
    i)
      GO_ACTION=install
      ;;
    *)
      showUsage
      ;;
  esac
done

## remove options
shift $((OPTIND - 1))

## exiting?
if [ $should_exit == true ]; then
  exit 0
fi

## check release option
if [ $RELEASE_BUILD -eq 1 ]; then
  ## clean dist directory
  rm -rf ./dist/

  ## call gox to build our binaries
  CGO_ENABLED=0 gox \
    -osarch="linux/amd64 darwin/amd64 freebsd/amd64 openbsd/amd64 windows/amd64 windows/386" \
    -ldflags="-X main.appVersion=${RELEASE_VERSION} -s -w" \
    -output="./dist/${BUILD_NAME}-${RELEASE_VERSION}-{{.Arch}}-{{.OS}}/${BUILD_NAME}-${RELEASE_VERSION}" \
    > /dev/null >&1

  ## gox return
  RETURN_VALUE=$?

else

  ## build it
  CGO_ENABLED=0 go ${GO_ACTION} \
    -ldflags="-X main.appVersion=${RELEASE_VERSION} -s -w" \
    > /dev/null >&1

  ## go build return
  RETURN_VALUE=$?
fi

## check build status
if [ ${RETURN_VALUE} -ne 0 ]; then
  printf "\nError during build!\n"
  exit ${RETURN_VALUE}
fi

## check release option
if [ $RELEASE_BUILD -eq 1 ]; then
  ## package binaries
  printf "Packaging ... "

  ## package files
  pushd ./dist/ || exit 1 > /dev/null >&1
  find . -maxdepth 1 -type d -name \*-\* \
    -exec tar -czf {}.tar.gz {} \; \
    -exec zip -m -r {}.zip {} \; > /dev/null >&1

  ## all done
  printf "done.\nRelease files may be found in the ./dist/ directory.\n"
else
  ## all done
  printf "done.\nUsage: %s -h\n" "${BUILD_NAME}"
fi

## exit same as build
exit ${RETURN_VALUE}
