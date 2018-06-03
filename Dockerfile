FROM golang:latest
WORKDIR /go/src/

ENV GOBIN=/go/wowbot
ENV GOPATH=/go

COPY . .

RUN go get github.com/OisinA/WowBot

WORKDIR /go/src/github.com/OisinA/WowBot

RUN go install

ENV PATH=/go/wowbot:${PATH}

WORKDIR /go/wowbot

EXPOSE 8080

CMD ["WowBot"]