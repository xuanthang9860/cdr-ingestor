FROM golang:1.24-alpine AS builder

WORKDIR /src

# Copy source
COPY . .
# Initialize Go modules and build binary
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o cdr-ingestor cmd/server/main.go

## Stage 2: Minimal runner
FROM alpine:3.19

# Install CA certificates if your binary needs to call HTTPS APIs
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy only the binary
COPY --from=builder /src/cdr-ingestor .

# Run binary
CMD ["./cdr-ingestor"]
# docker build -t cdr-ingestor .
# docker run -it --rm --network host --name test cdr-ingestor