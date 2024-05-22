FROM alpine:latest

MAINTAINER Idan, <idankob@gmail.com>

RUN apk add --no-cache --update curl ca-certificates openssl \
    && adduser --disabled-password --home /home/container container

USER container
ENV  USER=container HOME=/home/container

COPY go.mod go.sum /app/
COPY luckperms-notifier /app/

WORKDIR /app/

CMD ["./luckperms-notifier"]