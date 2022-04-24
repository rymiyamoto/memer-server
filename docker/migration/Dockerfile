ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine

RUN apk --no-cache add \
  git \
  openssh \
  gcc \
  musl-dev \
  tzdata;

ENV GOPATH="/go" \
  GO_ENV=test \
  DB_USER=root \
  DB_PASSWORD=test \
  DB_ADDRESS=rdb \
  DB_SCHEMA=memer

RUN set -uex \
  mkdir -p $GOPATH/src/github.com/rymiyamoto ; \
  go install github.com/rubenv/sql-migrate/...@latest ;
