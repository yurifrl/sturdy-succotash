# Build stage
FROM golang:alpine AS builder

WORKDIR /usr/local/app

RUN apk add --no-cache --update ca-certificates git
RUN update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /entrypoint main.go

# FROM scratch
FROM alpine
COPY --from=builder /entrypoint /entrypoint
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/entrypoint"]
