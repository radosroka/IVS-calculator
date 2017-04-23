pwd=$(shell pwd)
goexport=export GOPATH=$(pwd);

all: build documentation

gtkbuild:
	$(goexport) go get github.com/radosroka/go-gtk/gtk || true
	$(goexport) go get github.com/pkg/profile || true

build: gtkbuild
	mkdir -p bin
	$(goexport) cd bin/ && go build ../src/ivs-calc.go
	$(goexport) cd bin/ && go build -gcflags '-N -l' ../src/proff.go

run:
	bin/ivs-calc

pprof:
	cat profiling/vstup.txt | nice -n 19 bin/proff 1000
	go tool pprof bin/proff profiling/cpu.pprof

documentation:
	mkdir -p doc
	$(goexport) godoc -url=/pkg/calculator > doc/calc.html
	$(goexport) godoc -url=/pkg/gui > doc/gui.html

clean:
	go clean
	rm -rf bin/* profiling/cpu.proff doc/*

clean-all: clean
	rm -rf pkg src/github.com
	
