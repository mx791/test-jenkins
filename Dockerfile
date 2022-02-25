FROM golang:1.12-alpine AS build_base

COPY . /usr/src/app

WORKDIR /usr/src/app
RUN go build etl.go
CMD ["/usr/src/app/etl"]