FROM golang:latest

LABEL maintainer="rohitha.k95@gmail.com"


RUN mkdir /github.com/tb/task-logger
ADD . /github.com/tb/task-logger
WORKDIR /github.com/tb/task-logger

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8086

RUN go build -o main .

CMD ["/github.com/tb/cache-manager/backend/golang/taskapp/server.go"]

