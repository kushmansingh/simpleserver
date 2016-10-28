FROM golang:1.7-alpine

EXPOSE 80
COPY main.go /go/src/simple-server/
WORKDIR /go/src/simple-server/
RUN go install

ENTRYPOINT /go/bin/simple-server
