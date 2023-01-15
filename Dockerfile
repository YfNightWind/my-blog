FROM golang:1.19

MAINTAINER Maintainer

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /my-blog
COPY . /my-blog

RUN go build .

EXPOSE 3000

ENTRYPOINT ["./my-blog"]

