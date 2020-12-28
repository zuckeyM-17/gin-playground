FROM golang:latest

COPY src/app /go/src/app

WORKDIR /go/src/app/

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]