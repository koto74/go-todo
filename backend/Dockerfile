FROM golang:latest

RUN mkdir /go/src/app

WORKDIR /go/src/app

COPY ./backend/src/* /go/src/app

RUN apt-get update \
    && go mod init backend \
    && go get github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql

CMD "go" "run" "main.go"
