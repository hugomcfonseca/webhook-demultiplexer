GOOS = linux
ARCHS = amd64

all: build install

build: deps
    go build

deps:
    go get -d -v -t ./...

install:
