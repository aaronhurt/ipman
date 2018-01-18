#!/usr/bin/env bash
#
## package declarations
BUILD_NAME="ipman"
RELEASE_VERSION="0.1"
RELEASE_BUILD=0

## simple usage example
showUsage() {
	printf "Usage: $0 [-u|-i|-d]
	-u    Update vendor directory from glide.yaml using 'glide up' and build
	-i    Install vendor directory from glide.lock using 'glide install' and build
	-d    Remove existing glide.lock and vendor directory and exit
	-r    Build and package release binaries\n\n"
	exit 0
}

## install glide if needed
ensureGlide() {
	which glide > /dev/null 2>&1
	if [ $? -ne 0 ]; then
		printf "Installing glide ... "
		go get github.com/Masterminds/glide
	fi
}

## install gox if needed
ensureGox() {
	which gox > /dev/null 2>&1
	if [ $? -ne 0 ]; then
		printf "Installing gox ... "
		go get github.com/mitchellh/gox
	fi
}

## exit toggle
should_exit=false

## read options
while getopts ":uidcr" opt; do
	case $opt in
		u)
			ensureGlide
			printf "Updating vendor directory ... "
			glide -q up > /dev/null 2>&1
		;;
		i)
			ensureGlide
			printf "Installing from glide.lock ... "
			glide -q install > /dev/null 2>&1
		;;
		d)
			printf "Removing binary, glide.lock and vendor directory ... "
			rm -rf "${BUILD_NAME}" glide.lock vendor
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
		*)
			showUsage
		;;
	esac
done

## remove options
shift $((OPTIND-1))

## exiting?
if [ $should_exit == true ]; then
	exit 0
fi

## check release option
if [ $RELEASE_BUILD -eq 1 ]; then
	## clean dist directory
	rm -rf ./dist/

	## build release
	printf "Building release ... "

	## call gox to build our binaries
	CGO_ENABLED=0 gox \
	-osarch="linux/amd64 darwin/amd64 freebsd/amd64 openbsd/amd64 windows/amd64 windows/386" \
	-ldflags="-X main.appVersion=${RELEASE_VERSION} -s -w" \
	-output="./dist/${BUILD_NAME}-${RELEASE_VERSION}-{{.Arch}}-{{.OS}}/${BUILD_NAME}-${RELEASE_VERSION}" \
	> /dev/null >&1

	## gox return
	RETVAL=$?

else

	## build binaries
	printf "Building ... "

	## build it
	CGO_ENABLED=0 go build -o "${BUILD_NAME}" \
	-ldflags="-X main.appVersion=${RELEASE_VERSION} -s -w" \
	> /dev/null >&1

	## go build return
	RETVAL=$?
fi

## check build status
if [ $RETVAL -ne 0 ]; then
	printf "\nError during build!\n"
	exit $RETVAL
fi

## check release option
if [ $RELEASE_BUILD -eq 1 ]; then
	## package binaries
	printf "Packaging ... "

	## package files
	pushd ./dist/ > /dev/null >&1
	find . -maxdepth 1 -type d -name \*-\* \
	-exec tar -czf {}.tar.gz {} > /dev/null >&1 \; \
	-exec zip -m -r {}.zip {} > /dev/null >&1 \;

	## all done
	printf "done.\nRelease files may be found in the ./dist/ directory.\n"
else
	## all done
	printf "done.\nUsage: ./${BUILD_NAME} -h\n"
fi

## exit same as build
exit $RETVAL
