FROM golang:latest

RUN mkdir -p /app/backend/src

WORKDIR /app/backend/src

COPY src/* /app/backend/src

RUN apt-get update \
    && go get github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql

CMD "go" "run" "main.go"
