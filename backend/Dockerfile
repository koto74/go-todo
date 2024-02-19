FROM golang:latest

RUN mkdir /src/backend

WORKDIR /src/backend

COPY go.mod .
COPY go.sum .

RUN apt-get update \
    && go mod download \
    && go get github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql

COPY . .

ENV CGO_ENABLED=0

RUN  GOOS=linux go build -a -installsuffix cgo -o main .
