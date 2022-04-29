ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine

RUN apk add --no-cache \
    gcc \
    musl-dev \
    ca-certificates \
    bash \
    git ;

ENV GO_ENV=test \
    DB_USER=root \
    DB_PASSWORD=test \
    DB_ADDRESS=rdb \
    DB_SCHEMA=memer

ENV GOPATH="/go"
RUN mkdir -p $GOPATH/src/github.com/rymiyamoto

COPY ./run /usr/local/bin/run
RUN chmod +x /usr/local/bin/run

CMD ["run"]
