root?=$(abspath .)
build?=$(root)/bin
install?=$(GOPATH)/bin
src?=$(root)/cmd
ldflags?='-s -w'
formatter?=$(shell which gofumpt)

all: generate build install

default: all

generate: $(formatter)
	go generate ./...
	$(formatter) -e -w .

build:
	mkdir -p $(build)
	go build -ldflags=$(ldflags) -o $(build)/stdout $(src)/stdout 
	go build -ldflags=$(ldflags) -o $(build)/stderr $(src)/stderr 

install:
	ln -fs $(build)/stdout $(install)/stdout
	ln -fs $(build)/stderr $(install)/stderr
