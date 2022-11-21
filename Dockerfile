FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app

FROM alpine:latest
COPY --from=builder /app /app
WORKDIR /app
EXPOSE 8081
CMD ["/app/app"]