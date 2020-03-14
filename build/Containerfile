FROM golang:alpine AS builder

MAINTAINER 0xO1

RUN apk --update add ca-certificates git

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR /go/src/github.com/0x0I/aws_ec2_exporter

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

ADD src/ /go/src/github.com/0x0I/aws_ec2_exporter/

RUN GOPATH=/go go get && GOPATH=/go go build -o /bin/aws_ec2_exporter

FROM alpine:latest

COPY --from=builder /bin/aws_ec2_exporter /usr/bin/aws_ec2_exporter

ENTRYPOINT [ "aws_ec2_exporter" ]
