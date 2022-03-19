#!/bin/bash
FROM golang:1.16-alpine

WORKDIR /usr/local/go/src/gza

COPY ./go.mod /usr/local/go/src/gza
COPY ./go.sum /usr/local/go/src/gza
RUN go mod download

COPY . /usr/local/go/src/gza
RUN go build /usr/local/go/src/gza/server.go
RUN chmod +x /usr/local/go/src/gza/server

CMD ["/usr/local/go/src/gza/server"]

EXPOSE 8080