FROM golang:1.13-alpine3.11
LABEL maintainer="teruya ono"
WORKDIR /go/src/app
RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh
EXPOSE 8080
CMD [ "fresh" ]
