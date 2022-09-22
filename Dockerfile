FROM golang:latest

WORKDIR /app

COPY *.mod ./
COPY *.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o app ./cmd/main.go
