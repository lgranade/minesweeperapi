#***********************************************

FROM golang:1.15.3-alpine3.12 as builder

WORKDIR /go/src/github.com/lgranade/minesweeperapi

COPY src ./

RUN go build -o /go/bin/minesweeperapi

#***********************************************

FROM golang:1.15.3-alpine3.12 as runner

WORKDIR /opt/mineswpeerapi/

COPY --from=builder \
  /go/bin/minesweeperapi \
  .

EXPOSE 8080

CMD ["/opt/mineswpeerapi/mineswpeerapi"]

#***********************************************