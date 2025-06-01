FROM golang:1.24 as builder
WORKDIR /app
COPY . .

RUN go mod download
WORKDIR /app/cmd
RUN go build -o /app/app .


FROM scratch
WORKDIR /app
COPY --from=builder /app/app /app/
COPY --from=builder /app/.env /app/

EXPOSE 8080

CMD ["./main"]