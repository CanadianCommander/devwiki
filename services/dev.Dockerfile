FROM golang:1.20-alpine AS builder

ARG SERVICE

WORKDIR /var/src
COPY . .

RUN go build -o ./$SERVICE/server ./$SERVICE/cmd/$SERVICE

WORKDIR /var/src/$SERVICE
ENTRYPOINT ["./server"]