FROM golang:1.17-alpine

ADD . /code
WORKDIR /code

# Compile all cmd/* for linux
RUN ls -d -- ./cmd/*/ | cut -f3 -s -d '/' | CGO_ENABLED=0 GOOS=linux xargs -I {} go build \
    -o /bin/{} ./cmd/{}/...
