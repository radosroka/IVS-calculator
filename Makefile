pwd=$(shell pwd)
goexport=export GOPATH=$(pwd);
RPM_DIRS = --define "_sourcedir `pwd`" \
       --define "_rpmdir `pwd`" \
       --define "_specdir `pwd`" \
       --define "_builddir `pwd`" \
       --define "_srcrpmdir `pwd`"

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

deb-build: build
	mkdir -p deb/ivs-calc-1.0-1/usr/local/bin
	cp bin/ivs-calc	deb/ivs-calc-1.0-1/usr/local/bin/
	cd deb && dpkg-deb --build ivs-calc-1.0-1
	mv deb/ivs-calc-1.0-1.deb ./

deb-install: deb-build
	sudo apt -f install ivs-calc-1.0-1.deb

deb-uninstall: 
	sudo apt remove ivs-calc

rpm-build: pack
	rpmbuild $(RPM_DIRS) -ba rpm/golang-ivs-calculator.spec

documentation:
	mkdir -p doc
	$(goexport) godoc -url=/pkg/calculator > doc/calc.html
	$(goexport) godoc -url=/pkg/gui > doc/gui.html

pack:
	git archive --format=tar --prefix=IVS-calculator-1.1/ v1.1 | gzip > IVS-calculator-1.1.tar.gz

clean:
	go clean
	rm -rf bin/* profiling/cpu.proff deb/ivs-calc-1.0-1/usr/local/bin/ivs-calc

clean-all: clean
	rm -rf pkg src/github.com *.deb *.rpm IVS-calculator-1.1.tar.gz
	
