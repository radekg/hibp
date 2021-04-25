FROM golang:1.16.3-alpine3.13 as builder
RUN apk add alpine-sdk ca-certificates

WORKDIR /go/src/github.com/radekg/hibp
COPY . .

RUN make build

FROM alpine:3.13
COPY --from=builder /etc/ca-certificates /etc/ca-certificates
COPY --from=builder /go/src/github.com/radekg/hibp/hibp-linux-amd64 /opt/hibp/bin/hibp-linux-amd64
ENTRYPOINT ["/opt/hibp/bin/hibp-linux-amd64"]
CMD ["--help"]