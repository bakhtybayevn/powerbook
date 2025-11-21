# ============================
# 1) Builder Stage
# ============================
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates
WORKDIR /src

# Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build optimized binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -o /out/powerbook ./cmd/powerbook

# ============================
# 2) Production Image (distroless)
# ============================
FROM gcr.io/distroless/static:nonroot

# Copy compiled binary
COPY --from=builder /out/powerbook /usr/local/bin/powerbook

# Copy configuration folder
COPY --from=builder /src/internal/config /app/internal/config

# Declare port
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/usr/local/bin/powerbook"]
