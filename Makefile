SHELL := /bin/bash

default: all

all: build

deps:
	go get -d -v -u github.com/micro/go-grpc
	go get -d -v -u github.com/micro/kubernetes/cmd/micro
	go get -d -v -u github.com/micro/protobuf/{proto,protoc-gen-go}

build: deps
	go build -o mymicro

.PTHONY: all deps build