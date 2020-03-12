
FROM golang:alpine

MAINTAINER 0xO1

COPY . /go/src/github.com/0xO1/aws_ec2_exporter

WORKDIR /go/src/github.com/0xO1/aws_ec2_exporter

RUN GOPATH=/go go get && GOPATH=/go go build -o /bin/aws_ec2_exporter

ENTRYPOINT [ "/bin/aws_ec2_exporter" ]
