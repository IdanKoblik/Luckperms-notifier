FROM golang:1.22.3-alpine3.19

MAINTAINER Idan, <idankob@gmail.com>

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

USER container
ENV  USER=container HOME=/home/container

COPY go.mod go.sum /app/
RUN go mod download

ADD . /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o /luckperms-notifier

CMD ["/bin/bash", "/luckperms-notifier"]

