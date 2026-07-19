# ---------- Build Stage ----------
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o devops-dashboard ./cmd/server

# ---------- Runtime Stage ----------
FROM alpine:3.22

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/devops-dashboard .
COPY --from=builder /app/internal ./internal

EXPOSE 8080

CMD ["./devops-dashboard"]
