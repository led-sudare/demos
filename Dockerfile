# BUILD
FROM golang:1.12-alpine3.10

RUN apk add --no-cache \
    ca-certificates \
    make \
    git \
    gcc \
    g++ \
    czmq-dev \
    libzmq \
    libsodium

ENV GO111MODULE=on
WORKDIR /work/
EXPOSE 5003

ENTRYPOINT ["./docker-entrypoint.sh"]

