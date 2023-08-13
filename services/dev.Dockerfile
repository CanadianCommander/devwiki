FROM golang:1.21-alpine AS builder

ARG SERVICE

RUN cd /root/ && go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /var/src
COPY . .

ENV SERVICE=$SERVICE
WORKDIR /var/src/$SERVICE
ENTRYPOINT ["sh", "-c", "CompileDaemon -directory='.'  -directory='../' -log-prefix=false -build=\"go build -o ./server ./cmd/${SERVICE}\" -command=./server"]