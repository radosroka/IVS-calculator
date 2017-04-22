pwd=$(shell pwd)
goexport=export GOPATH=$(pwd);

all: build

gtkbuild:
	$(goexport) go get github.com/mattn/go-gtk/gtk || true

build: gtkbuild
	mkdir -p bin
	$(goexport) cd bin/ && go build ../src/gui.go
	$(goexport) cd bin/ && go build ../src/proff.go

run:
	bin/gui

clean:
	go clean
	rm -rf bin/*

clean-all: clean
	rm -rf pkg src/github.com
	
