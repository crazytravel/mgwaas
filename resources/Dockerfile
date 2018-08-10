FROM golang:1.10.3-alpine

WORKDIR /go/src/mgwaas
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["mgwaas"]