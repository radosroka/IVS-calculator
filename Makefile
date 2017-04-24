# IVS-calculator
# Copyright (C) 2017	Radovan Sroka <xsroka00@stud.fit.vutbr.cz>
# 						Tomáš Sýkora <xsykor25@stud.fit.vutbr.cz>
#						Michal Cyprian <xcypri01@stud.fit.vutbr.cz>
#						Jan Mochnak <xmochn00@stud.fit.vutbr.cz>
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.


pwd=$(shell pwd)
goexport=export GOPATH=$(pwd);
RPM_DIRS = --define "_sourcedir `pwd`" \
       --define "_rpmdir `pwd`" \
       --define "_specdir `pwd`" \
       --define "_builddir `pwd`" \
       --define "_srcrpmdir `pwd`"

files=$(shell ls)
files += .git


all: build doc

gtkbuild:
	$(goexport) go get github.com/radosroka/go-gtk/gtk || true
	$(goexport) go get github.com/pkg/profile || true

build: gtkbuild
	mkdir -p bin
	$(goexport) cd bin/ && go build ../src/ivs-calc.go
	$(goexport) cd bin/ && go build -gcflags '-N -l' ../src/proff.go

run:
	bin/ivs-calc

test:
	$(goexport) go test calculator

pprof:
	cat profiling/vstup.txt | nice -n 19 bin/proff 1000
	go tool pprof bin/proff profiling/cpu.pprof

deb-build: build
	mkdir -p deb/ivs-calc-1.0-1/usr/local/bin
	cp bin/ivs-calc	deb/ivs-calc-1.0-1/usr/local/bin/
	cd deb && dpkg-deb --build ivs-calc-1.0-1
	mv deb/ivs-calc-1.0-1.deb ./

deb-install: deb-build
	sudo apt -f install ./ivs-calc-1.0-1.deb

deb-uninstall:
	sudo apt remove ivs-calc

rpm-build: rpm-archive
	sudo dnf -y install golang golang-godoc gtk2-devel pango-devel
	rpmbuild $(RPM_DIRS) -ba rpm/golang-ivs-calculator.spec

rpm-install: rpm-build
	sudo dnf install -y x86_64/*.rpm

rpm-uninstall:
	sudo dnf remove -y golang-ivs-calculator

rpm-archive:
	git archive --format=tar --prefix=IVS-calculator-1.1/ v1.1 | gzip > IVS-calculator-1.1.tar.gz


docker-build: rpm-build
	docker-compose build

docker-run:
	sudo setenforce 0 || true
	sudo xhost local:root
	sudo docker-compose up

doc:
	mkdir -p doc
	$(goexport) godoc -url=/pkg/calculator | wkhtmltopdf - doc/doc-calc.pdf 2>/dev/null || true
	$(goexport) godoc -url=/pkg/gui | wkhtmltopdf - doc/doc-gui.pdf 2>/dev/null || true
	pdfunite doc/* dokumentace.pdf

pack: all deb-build
	mkdir -p xsroka00_xsykor25_xcypri01_xmochn00
	mkdir -p xsroka00_xsykor25_xcypri01_xmochn00/repo
	cp -r $(files) xsroka00_xsykor25_xcypri01_xmochn00/repo/
	cp dokumentace.pdf doc/dokumentace.pdf
	cp -r doc/ xsroka00_xsykor25_xcypri01_xmochn00/doc
	mkdir -p xsroka00_xsykor25_xcypri01_xmochn00/install
	cp *.deb xsroka00_xsykor25_xcypri01_xmochn00/install/
	zip -r xsroka00_xsykor25_xcypri01_xmochn00.zip xsroka00_xsykor25_xcypri01_xmochn00
	rm -rf xsroka00_xsykor25_xcypri01_xmochn00 doc/dokumentace.pdf

clean:
	go clean
	rm -rf bin profiling/cpu.proff deb/ivs-calc-1.0-1/usr/local/bin/ivs-calc doc/
	rm -rf pkg x86_64 src/github.com *.deb *.rpm IVS-calculator-1.1.tar.gz
	rm -rf dokumentace.pdf xsroka00_xsykor25_xcypri01_xmochn00/ xsroka00_xsykor25_xcypri01_xmochn00.zip
