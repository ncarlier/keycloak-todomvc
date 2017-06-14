FROM golang:1.7.3

RUN apt-get update
RUN apt-get install jq

RUN mkdir -p /go/src/todo

WORKDIR /go/src/todo

ARG TODOMVC_CLIENT_SECRET
ENV TODOMVC_CLIENT_SECRET $TODOMVC_CLIENT_SECRET

ARG API_ENDPOINT
ENV API_ENDPOINT $API_ENDPOINT

ARG AUTH_ENDPOINT
ENV AUTH_ENDPOINT $AUTH_ENDPOINT

RUN go get -v github.com/urfave/cli

ADD src/todo .
ADD todomvc-cli.go todomvc-cli.go

RUN env GOOS=linux GOARCH=amd64 go build -v -o /usr/local/sbin/todomvc

ENTRYPOINT ["/bin/bash"]