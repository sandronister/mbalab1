FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN go build -o /app/app .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./app"]