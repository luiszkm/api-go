FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . /app

# RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server ./cmd/server/main.go

RUN go mod tidy

# ENTRYPOINT [ "/server" ] 

