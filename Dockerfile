FROM golang:1.15-alpine

WORKDIR /go/awesomeProject

COPY src/app ./src/app
COPY go.mod .

RUN apk add --no-cache git \
  && go get github.com/oxequa/realize


WORKDIR /go/awesomeProject/src/app

RUN go build -o app

CMD ["/go/awesomeProject/src/app"]
