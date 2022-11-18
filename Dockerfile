FROM golang:1.18.7-alpine3.16 as builder

# Add Maintainer Info
LABEL maintainer="Atsuhiro Endo <atcheri@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN ls -la

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o http ./cmd/main.go

######## Start a new stage from scratch #######
FROM scratch

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/http .
COPY --from=builder /app/internal/core/adapters/right/repositories/inMemory/languages.json ./internal/core/adapters/right/repositories/inMemory/languages.json

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./http"]