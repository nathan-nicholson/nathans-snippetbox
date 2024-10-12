FROM --platform=$BUILDPLATFORM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./ui ./ui

# Build the Go app
RUN env ${opts} go build -v -o /app/snippetbox ./cmd/web

# Use a minimal image to run the binary
FROM --platform=$BUILDPLATFORM ubuntu:24.10

# Copy the pre-built binary from the dist directory to the container
COPY --from=builder /app/snippetbox /app/snippetbox

# Copy the UI files
COPY --from=builder /app/ui /ui

# Expose the port that the service will run on
EXPOSE 4000

# Command to run the binary
CMD ["/app/snippetbox"]
