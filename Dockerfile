FROM golang:1.14.2-alpine3.11

WORKDIR /go
ADD . /go
RUN apk update && \
    apk add git vim
CMD go run main.go