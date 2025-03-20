FROM golang:1.24.1-alpine3.21 AS builder

WORKDIR /app

# Build the binary
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main cmd/*.go 

# Build a small image
FROM golang:1.24.1-alpine3.21

COPY --from=builder /app/main /usr/local/bin/

EXPOSE 9000

CMD ["main"]
