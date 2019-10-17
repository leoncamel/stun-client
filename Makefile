
all:build

clean:
	rm -rf dist

build:
	go build

release:
	goreleaser --snapshot --skip-publish --rm-dist

