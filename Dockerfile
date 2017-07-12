FROM golang:1.8-alpine

MAINTAINER sgordillogallardo@gmail.com
ENV NAME vizzuality
ENV PROJECT poc-bigquery-geo

RUN apk update && apk upgrade && \
    apk add --no-cache --update bash git

# Install dependencies
RUN go get github.com/codegangsta/gin

# Code
ADD . /go/src/github.com/$NAME/$PROJECT

# Install & Run
WORKDIR /go/src/github.com/$NAME/$PROJECT
RUN go-wrapper download
RUN go-wrapper install
EXPOSE 3050

ENTRYPOINT ["./entrypoint.sh"]
