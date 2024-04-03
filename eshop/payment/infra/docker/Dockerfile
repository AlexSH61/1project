FROM golang:1.22 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app payment/app/services/payment-api/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]