FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -o main .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]

EXPOSE 8080
