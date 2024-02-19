FROM golang:latest

RUN mkdir -p /app/backend

WORKDIR /app/backend

#RUN apt-get update \
    #&& go mod init backend \
    #&& go get github.com/gin-gonic/gin \
    #&& go get github.com/jinzhu/gorm \
    #&& go get github.com/go-sql-driver/mysql

COPY src .

CMD "go" "run" "main.go"

