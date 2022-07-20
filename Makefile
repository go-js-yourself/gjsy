.PHONY: test build docker-image

default: test build

docker-image:
	DOCKER_BUILDKIT=1 docker build -t gjsy .

test: docker-image
	docker run -v $$PWD:/build --rm gjsy go test ./...

build: docker-image
	docker run --name gjsy-build -v $$PWD:/build -e CGO_ENABLED=0 gjsy \
		go install ./... && ls -alsh
	docker cp gjsy-build:/go/bin/repl .
	docker rm gjsy-build

