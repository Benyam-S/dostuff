FROM golang:1.15.7-alpine3.13

ADD . /dostuff
WORKDIR /dostuff

RUN apk add git
RUN go mod download
RUN go build -o main .

EXPOSE 8080

CMD  ["/dostuff/main"]