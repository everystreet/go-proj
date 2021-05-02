FROM golang:latest AS proj

ARG PROJ_VERSION=6.2.1

RUN apt-get -y update

RUN apt-get install -y automake libtool sqlite3 libsqlite3-dev unzip

RUN git clone --branch ${PROJ_VERSION} https://github.com/OSGeo/PROJ.git

WORKDIR /go/PROJ

RUN ./autogen.sh

RUN ./configure

RUN make

RUN make install

RUN mkdir -p /usr/local/share/proj; \
    curl -LOs https://download.osgeo.org/proj/proj-datumgrid-1.8.zip && unzip -j -u -o proj-datumgrid-1.8.zip -d /usr/local/share/proj; \
    curl -LOs https://download.osgeo.org/proj/proj-datumgrid-europe-1.2.zip && unzip -j -u -o proj-datumgrid-europe-1.2.zip -d /usr/local/share/proj; \
    curl -LOs https://download.osgeo.org/proj/proj-datumgrid-oceania-1.0.zip && unzip -j -u -o proj-datumgrid-oceania-1.0.zip -d /usr/local/share/proj; \
    curl -LOs https://download.osgeo.org/proj/proj-datumgrid-world-1.0.zip && unzip -j -u -o proj-datumgrid-world-1.0.zip -d /usr/local/share/proj; \
    curl -LOs https://download.osgeo.org/proj/proj-datumgrid-north-america-1.2.zip && unzip -j -u -o proj-datumgrid-north-america-1.2.zip -d /usr/local/share/proj; \
    rm *.zip

RUN make check

ENV LD_LIBRARY_PATH=${LD_LIBRARY_PATH}:/usr/local/lib

FROM proj

WORKDIR /usr/src/go-proj

COPY . .

RUN go test ./... -v -count 1
