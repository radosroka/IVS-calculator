FROM fedora:latest

MAINTAINER Michal Cyprian <xcypri01@stud.fit.vutbr.cz>

COPY golang-ivs-calculator-1.0-1.fc26.x86_64.rpm /tmp

RUN dnf install -y /tmp/golang-ivs-calculator-1.0-1.fc26.x86_64.rpm && \
    dnf clean all

RUN useradd -u 1001 -r -g 0 -d ${HOME} -s /sbin/nologin \
    -c "Default Application User" default

USER 1001

CMD bin/ivs-calc
