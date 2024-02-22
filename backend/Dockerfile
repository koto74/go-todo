FROM golang:latest

RUN mkdir /backend

WORKDIR /backend

COPY src /backend

RUN apt-get update \
    && go get github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql

CMD "go" "run" "main.go"
