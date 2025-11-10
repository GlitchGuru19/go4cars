# Build stage
ARG GO_VERSION=1.25.4
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -v -o /run-app .

# Final stage (minimal image)
FROM debian:bookworm-slim

# Copy the binary from builder
COPY --from=builder /run-app /usr/local/bin/
RUN chmod +x /usr/local/bin/run-app

# Command to run the binary
CMD ["run-app"]
