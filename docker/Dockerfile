ARG GO_VERSION
FROM golang:$GO_VERSION AS builder

ARG BUILD_VERSION

LABEL maintainer="Uladzimir Strazhenko <strazhenkoandi@gmail.com>" \
      org.opencontainers.image.created="${BUILD_VERSION}" \
      org.opencontainers.image.title="Random Data API" \
      org.opencontainers.image.description="Random Data API" \
      org.opencontainers.image.source="https://github.com/uladzimirSTR/randomData_API"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd

FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    htop \
    iproute2 \
    util-linux \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/main
COPY --from=builder /app/templates /app/templates

EXPOSE 9909

CMD ["/app/main"]
