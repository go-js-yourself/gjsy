.PHONY: test build docker-image

default: test build

docker-image:
	DOCKER_BUILDKIT=1 docker build -t gjsy .

test: docker-image
	docker run -v $$PWD:/go/src/github.com/go-js-yourself/gjsy --rm gjsy go test ./...

build: docker-image
	docker run --name gjsy-build -v $$PWD:/go/src/github.com/go-js-yourself/gjsy -e CGO_ENABLED=0 gjsy \
		go install ./...
	docker cp gjsy-build:/go/bin/repl .
	docker rm gjsy-build

