FROM golang:latest

COPY src/app /go/src/app

WORKDIR /go/src/app/

RUN go get -u github.com/gin-gonic/gin

CMD ["go", "run", "main.go"]