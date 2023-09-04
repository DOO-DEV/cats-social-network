FROM golang:1.18-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/tinrab/meower

COPY go.mod ./
COPY go.sum  ./
COPY util util
COPY event event
COPY db db
COPY search search
COPY schema schema
COPY service/meow meow-service
COPY service/query query-service
COPY service/pusher pusher-service

RUN go mod download

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .