.PHONY: build release check clean distclean

build:
	@build/build.sh -i

release:
	@build/build.sh -ir

check:
	@build/codeCheck.sh

clean:
	@build/build.sh -d

distclean:
	@build/build.sh -dc
