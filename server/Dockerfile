FROM golang:1.15.6-alpine3.12
LABEL maintainer="teruya ono"
WORKDIR /go/src/app
RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh
CMD [ "fresh" ]
