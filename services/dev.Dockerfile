FROM golang:1.20-alpine AS builder

ARG SERVICE

RUN cd /root/ && go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /var/src
COPY . .

# RUN go build -o ./$SERVICE/server ./$SERVICE/cmd/$SERVICE

ENV SERVICE=$SERVICE
WORKDIR /var/src/$SERVICE
ENTRYPOINT ["sh", "-c", "CompileDaemon -log-prefix=false -build=\"go build -o ./server ./cmd/${SERVICE}\" -command=./server"]