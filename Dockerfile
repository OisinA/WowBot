FROM golang:latest
WORKDIR /go/src/github.com/OisinA/WowBot

ENV GOBIN=/go/wowbot
ENV GOPATH=/go

COPY . .

RUN go get ./
RUN go build

WORKDIR /go/wowbot

EXPOSE 8080

CMD ["./WowBot"]