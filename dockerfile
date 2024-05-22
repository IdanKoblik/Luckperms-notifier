FROM golang:1.22.3-alpine3.19

MAINTAINER Idan, <idankob@gmail.com>

RUN apk add --no-cache --update curl ca-certificates openssl git tar bash sqlite fontconfig \
    && adduser --disabled-password --home /home/container container

USER container
ENV  USER=container HOME=/home/container

COPY go.mod go.sum /app/
COPY . /app/

USER root
RUN chown -R container:container /app

USER container

WORKDIR /app/
RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath 

RUN chmod +x luckperms-notifier

CMD ["./luckperms-notifier"]