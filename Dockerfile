FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o sarc

# Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sarc .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./sarc"]