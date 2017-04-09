pwd=$(shell pwd)
goexport=export GOPATH=$(pwd);
ldexport=export LD_LIBRARY_PATH=$(pwd)/src/github.com/visualfc/goqt/bin;

all: build

qtbuild: 
	$(goexport) go get -d github.com/visualfc/goqt || true
	cd $(pwd)/src/github.com/visualfc/goqt/qtdrv && qmake-qt4 "CONFIG+=release" || qmake "CONFIG+=release" && make
	cd $(pwd)/src/github.com/visualfc/goqt/tools/rcc && qmake-qt4 "CONFIG+=release" || qmake "CONFIG+=release" && make
	cd $(pwd)/src/github.com/visualfc/goqt/ui ; go install -v


build: qtbuild
	mkdir -p bin
	$(goexport) cd bin/ && go build ../src/main.go
	$(goexport) cd bin/ && go build ../src/example.go

run-example:
	$(ldexport) bin/example	

run-calc:
	$(ldexport) bin/main	

run: run-example run-calc



clean:
	go clean
	rm -rf bin/*

clean-all: clean
	rm -rf pkg src/github.com
	
