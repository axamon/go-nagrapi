FROM        frolvlad/alpine-glibc:latest
MAINTAINER  François ALLAIS <francois.allais@hotmail.com>

ADD go-nagrapi /usr/bin

RUN mkdir /data

WORKDIR /usr/bin

EXPOSE     5555
VOLUME     [ "/data" ]
CMD        [ "/usr/bin/go-nagrapi", "--s", "/data/status.dat" ]
