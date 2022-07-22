# Build stage
FROM golang:1.18.4-alpine3.16 as build

WORKDIR /go/bin/

COPY go.mod ./

COPY main.go ./

RUN go build -o ./depends

# Execution stage
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/depends /

ENTRYPOINT ["/depends"]