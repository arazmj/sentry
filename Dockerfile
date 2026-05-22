FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/sentry-server ./server \
    && CGO_ENABLED=0 GOOS=linux go build -o /out/sentry ./cmd/cli

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=builder /out/sentry-server /usr/local/bin/
COPY --from=builder /out/sentry /usr/local/bin/

EXPOSE 50051
ENTRYPOINT ["/usr/local/bin/sentry-server"]
