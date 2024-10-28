# Stage 1: Build the Go application
FROM golang:1.21 AS builder
WORKDIR /app

# Copy the entire project to the /app directory
COPY . .

# Change directory to main (where main.go is located)
WORKDIR /app/main

# Download dependencies and build the application
RUN go mod download
RUN go mod tidy
RUN go mod vendor
RUN go build -o payment-gateway .

# Stage 2: Create the final image
FROM gcr.io/distroless/base
COPY --from=builder /app/main/payment-gateway /app/payment-gateway
CMD ["/app/payment-gateway"]
