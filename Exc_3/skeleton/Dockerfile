  # todo
# ---------- builder ----------
FROM golang:1.25 AS builder
WORKDIR /app

# copy and cache modules first
COPY go.mod go.sum ./
RUN go mod download

# copy the whole project
COPY . .

# run the provided build script
RUN chmod +x scripts/build-application.sh && ./scripts/build-application.sh

# ---------- runtime ----------
FROM alpine:latest
RUN adduser -D appuser
WORKDIR /app

# copy the binary from the builder stage
COPY --from=builder /app/ordersystem /app/ordersystem

EXPOSE 3000
USER appuser
ENTRYPOINT ["./ordersystem"]
