FROM golang:alpine AS builder
WORKDIR /build
COPY /src .
RUN CGO_ENABLED=0 GOOS=linux go build -o temperatureapi *.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/temperatureapi .
EXPOSE 8081
CMD ["./temperatureapi"]