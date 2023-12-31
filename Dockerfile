FROM golang:1.21.0-alpine3.18 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/tinrab/meower

COPY go.mod go.sum ./
COPY util util
COPY event event
COPY db db
COPY search search
COPY schema schema
COPY meow-service meow-service
COPY query-service query-service
COPY pusher-service pusher-service

RUN GO111MODULE=on go mod download
RUN GO111MODULE=on go install ./...

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .