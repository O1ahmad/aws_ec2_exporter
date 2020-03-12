FROM golang:alpine AS builder

MAINTAINER 0xO1

ADD . /go/src/github.com/0x0I/aws_ec2_exporter/

WORKDIR /go/src/github.com/0x0I/aws_ec2_exporter/src

RUN apk --update add ca-certificates git \
 && GOPATH=/go go get \
 && GOPATH=/go go build -o /bin/aws_ec2_exporter

FROM alpine:latest

COPY --from=builder /bin/aws_ec2_exporter /usr/bin/aws_ec2_exporter

ENTRYPOINT [ "aws_ec2_exporter" ]
