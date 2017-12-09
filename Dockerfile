FROM alpine:latest
MAINTAINER Sergey Besedin
EXPOSE 80 443

RUN apk --no-cache add \
    ca-certificates \
    git \
    curl \
    go \
    gcc \
    libc-dev

RUN go get github.com/requilence/integram

COPY entrypoint.sh /
COPY main.go /

ENTRYPOINT /entrypoint.sh

CMD go run main.go
