FROM golang:latest AS proj

ARG PROJ_VERSION=8.2.0

RUN apt-get -y update

RUN apt-get install -y automake libtool sqlite3 libsqlite3-dev unzip libtiff-dev libcurl4-openssl-dev

RUN git clone --branch ${PROJ_VERSION} https://github.com/OSGeo/PROJ.git

WORKDIR /go/PROJ

RUN ./autogen.sh

RUN ./configure

RUN make

RUN make install

RUN make check

ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib


FROM proj

WORKDIR /usr/src/go-proj

COPY . .

RUN go test ./... -v -count 1
