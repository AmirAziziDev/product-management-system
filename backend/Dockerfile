# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the application
RUN mkdir -p build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/main .

# Final stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/build/main .

EXPOSE 8080

CMD ["./main"]