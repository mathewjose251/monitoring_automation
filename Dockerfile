FROM golang:1.14

WORKDIR /go/src/app

ENV GOBIN=/go/src/app

COPY /go/src/app/main.go .

RUN go get -d github.com/prometheus/client_golang/prometheus
RUN go build -v .

CMD ["./app"]