FROM golang:latest

COPY src/app /go/src/app

WORKDIR /go/src/app/

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/utrack/gin-csrf
RUN go get -u github.com/gin-contrib/sessions

CMD ["air", "-c", ".air.toml"]