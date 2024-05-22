FROM golang:1.22.3-alpine3.19

MAINTAINER Idan, <idankob@gmail.com>

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

USER container
ENV  USER=container HOME=/home/container

COPY go.mod go.sum /app/

COPY . /app/

WORKDIR /app
RUN go mod download

RUN go build -ldflags="-buildmode=exe -trimpath -buildid= -w -s -buildvcs=false"

CMD ["/bin/bash", "./luckperms-notifier"]

