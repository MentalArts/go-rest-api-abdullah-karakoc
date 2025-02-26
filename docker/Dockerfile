FROM golang:1.19 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
