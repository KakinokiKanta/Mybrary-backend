FROM golang:1.22.1-alpine3.18

COPY ./src/ /go/src/

WORKDIR /go/src/

RUN apk update \
&& apk add --no-cache git \
&& go mod download && go mod verify
# && go get github.com/gin-gonic/gin \
# && go get github.com/jinzhu/gorm \
# && go get github.com/go-sql-driver/mysql

# EXPOSE 8080

CMD ["go", "run", "main.go"]
