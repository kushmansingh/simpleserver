FROM golang:1.10

EXPOSE 80
WORKDIR /go/src/simple-server/
COPY main.go /go/src/simple-server/
RUN go install

ENTRYPOINT ["/go/bin/simple-server"]
