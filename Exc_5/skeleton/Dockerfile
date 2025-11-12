
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh

# Runner
FROM alpine:latest

RUN adduser -D -h /home/appuser appuser

RUN apk add --no-cache ca-certificates bash netcat-openbsd

WORKDIR /app

COPY --from=builder /app/ordersystem /app/ordersystem

RUN chown appuser:appuser /app/ordersystem && chmod +x /app/ordersystem

USER appuser

#This doesn’t open the port; it’s just for reference and tooling
EXPOSE 3000

#Defines the command to run when the container starts
ENTRYPOINT ["./ordersystem"]
