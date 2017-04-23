pwd=$(shell pwd)
goexport=export GOPATH=$(pwd);

all: build documentation

gtkbuild:
	$(goexport) go get github.com/mattn/go-gtk/gtk || true
	$(goexport) go get github.com/pkg/profile || true

build: gtkbuild
	mkdir -p bin
	$(goexport) cd bin/ && go build ../src/run-gui.go
	$(goexport) cd bin/ && go build ../src/proff.go

run:
	bin/run-gui

pprof:
	cat numbers | bin/proff 1000
	go tool pprof bin/proff cpu.pprof

documentation:
	mkdir -p doc
	$(goexport) godoc -url=/pkg/calculator > doc/calc.html
	$(goexport) godoc -url=/pkg/gui > doc/gui.html

clean:
	go clean
	rm -rf bin/* doc/*

clean-all: clean
	rm -rf pkg src/github.com
	
