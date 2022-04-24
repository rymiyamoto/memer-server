ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine

RUN apk --no-cache add \
  git \
  openssh \
  gcc \
  musl-dev \
  tzdata;

ENV GOPATH="/go" \
  GO_ENV=test

RUN set -uex \
  mkdir -p $GOPATH/src/github.com/rymiyamoto ; \
  go get -v github.com/rubenv/sql-migrate/... ;
