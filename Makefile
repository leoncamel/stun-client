
all:build

clean:
	rm -rf dist

build:
	goreleaser --snapshot --skip-publish --rm-dist

