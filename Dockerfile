FROM golang:1.17.6-buster
WORKDIR /api/src
COPY ./src .
RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/gin-gonic/gin
CMD ["air", "-c", ".air.toml"]