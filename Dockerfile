# golang:1.23.1-bookworm
FROM golang@sha256:1a5326b07cbab12f4fd7800425f2cf25ff2bd62c404ef41b56cb99669a710a83 AS builder

WORKDIR /app

COPY . .

RUN go build -ldflags="-w -s" -o ./depends

# alpine:3.20.3
FROM alpine@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

COPY --from=builder /app/depends /

ENTRYPOINT ["/depends"]