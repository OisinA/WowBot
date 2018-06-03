FROM golang:latest
WORKDIR /go/src/github.com/OisinA/WowBot

ENV GOBIN=/go/WowBot
ENV GOPATH=/go

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /go/WowBot

CMD ["WowBot"]