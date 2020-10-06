FROM golang:latest

LABEL maintainer="rohitha.k95@gmail.com"

WORKDIR /github.com/tb/task-logger

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8086

RUN go build -o main .

CMD ["./"]

FROM golang:1.9.2
/github.com/tb/task-logger
RUN mkdir /github.com/tb/rohitha/dockerexampleout
ADD . /github.com/tb/rohitha/dockerexampleout
WORKDIR /github.com/tb/rohitha/dockerexampleout
RUN go build -o main .
CMD ["/github.com/tb/rohitha/dockerexampleout/main"]
