FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o emis cmd/main.go

FROM alpine:3.17.3

WORKDIR /app

COPY --from=builder /app/emis .

EXPOSE 8080

CMD ["./emis"]