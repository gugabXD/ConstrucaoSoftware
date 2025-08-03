FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# If go.mod and go.sum are missing, initialize a new Go module
COPY . .
RUN [ ! -f go.mod ] && go mod init sarc || true
RUN go mod tidy

# Install a specific version of swag
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

# Regenerate Swagger documentation
RUN rm -rf docs && swag init -g main.go -o docs

# Build the Go binary
RUN go build -o sarc

# Final image
FROM alpine:latest

WORKDIR /app

# Copy the built binary and Swagger docs from the builder stage
COPY --from=builder /app/sarc .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./sarc"]