FROM golang:1.17-alpine

ADD . /code
WORKDIR /code

# Install 0MQ
RUN apk update && \
    apk add --no-cache zeromq-dev libzmq

# Compile all cmd/* for linux
RUN ls -d ./cmd/*/ | cut -f3 -s -d '/' | CGO_ENABLED=1 GOOS=linux xargs -I {} go build \
    -o ./bin/{} ./cmd/{}/ ...
