.PHONY: test build docker-image repl run examples

default: test build

docker-image:
	@DOCKER_BUILDKIT=1 docker build -t gjsy .

test: docker-image
	docker run -v $$PWD:/go/src/github.com/go-js-yourself/gjsy --rm gjsy go test \
	  ./...

build/bin/gjsy: docker-image
	mkdir -p build
	docker run --rm \
	           -v $$PWD:/go/src/github.com/go-js-yourself/gjsy \
	           -v $$PWD/build:/go/bin \
	           -e CGO_ENABLED=0 \
	       gjsy go install ./...

build: build/bin/gjsy

repl: build/bin/gjsy
	./build/bin/gjsy repl

run: build/bin/gjsy
ifeq ($(FILE),)
	@echo "FILE variable is required"
else
	./build/bin/gjsy $(FILE)
endif

examples: build/bin/gjsy
	@for file in $(wildcard examples/*.js); do \
	  echo ====================================================================; \
	  echo File $$file ; \
	  echo ; \
	  cat $$file ; \
	  echo --------------------------------------------------------------------; \
	  echo Execution ; \
	  ./build/bin/gjsy $$file ; \
	  echo ; \
	done
