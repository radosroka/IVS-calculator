FROM fedora:latest

MAINTAINER Michal Cyprian <xcypri01@stud.fit.vutbr.cz>

RUN dnf install -y golang gtk2-devel libgda-devel libcanberra-gtk2 \
                   PackageKit-gtk3-module make && \
    dnf clean all

RUN mkdir -p /opt/app-root && \
    useradd -u 1001 -r -g 0 -d ${HOME} -s /sbin/nologin \
    -c "Default Application User" default && \
    chown -R 1001:0 /opt/app-root && chmod -R og+rwx /opt/app-root

COPY . /opt/app-root/

RUN cd /opt/app-root/ && pwd && ls -la
RUN  cd /opt/app-root && make

USER 1001

WORKDIR /opt/app-root

CMD bin/gui
