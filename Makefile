.PHONY: all build test docker tools dep gometalinter
all: lint test docker

build:
	# Used mostly to test compilation. doesn't actually build binaries.
	go build ./...

test:
	go test ./pkg/...

docker:
	docker build -t corentone/golang-boilerplate .

lint:
gometalinter --deadline 4m --sort=path --skip=pkg/mocks $$(glide nv)

tools: dep gometalinter

dep:
	go get -u github.com/golang/dep/cmd/dep

gometalinter:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install -u
