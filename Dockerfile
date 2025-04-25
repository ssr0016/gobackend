# Stage 1: Base image with Go and dependencies
FROM golang:1.22-alpine as base

# Install necessary packages
RUN apk update && \
    apk add --no-cache git alpine-sdk librdkafka-dev pkgconf

# Stage 2: Builder
FROM base as builder
WORKDIR /builder

# Copy and download Go dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build --ldflags "-extldflags -static" -tags musl -o /builder/app ./cmd/main.go

# Stage 3: Final image with only the binary
FROM alpine:latest

# Install bash (for your entrypoint script)
RUN apk add --no-cache --upgrade bash

# Set working directory
WORKDIR /app

# Copy the built binary and the entrypoint script from the builder stage
COPY --from=builder /builder/app ./app
COPY --from=builder /builder/docker-entrypoint.sh ./docker-entrypoint.sh

# Make the entrypoint script executable
RUN chmod +x ./docker-entrypoint.sh

# Create necessary directories (if needed for your app)
RUN mkdir -p ./public/img/

# Run the entrypoint script
ENTRYPOINT ["./docker-entrypoint.sh"]