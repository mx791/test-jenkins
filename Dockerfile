FROM golang:1.12-alpine AS build_base

COPY . /usr/src/app

WORKDIR /usr/src/app
RUN go build etl.go

FROM alpine
COPY --from=0 /usr/src/app /usr/src/app
WORKDIR /usr/src/app

CMD ["/usr/src/app/etl"]