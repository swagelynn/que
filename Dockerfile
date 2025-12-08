FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o app .

FROM alpine AS certs
RUN apk add --no-cache ca-certificates

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./app"]
