#########################################
# Build stage
#########################################
FROM golang:1.8 AS built-cli
MAINTAINER Nicolas Carlier <nicolas.carlier@worldline.com>

# Copy suources into the container
ADD . /go/src/todo

# Set working directory
WORKDIR /go/src/todo

# Build the binary
RUN go get github.com/urfave/cli && \
      GOOS=linux GOARCH=amd64 go build -o todomvc

#########################################
# Final stage
#########################################
#FROM alpine
FROM debian
MAINTAINER Nicolas Carlier <nicolas.carlier@worldline.com>

# Install JQ
#RUN apk add --no-cache jq
RUN apt update && apt install -y jq

# Fix lib dep
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# Install binary
COPY --from=built-cli /go/src/todo/todomvc /usr/local/bin/

