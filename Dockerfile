# /dev/Go/practice-api/Dockerfile
FROM golang:1.23

# Install Air for hot reload
RUN go install github.com/air-verse/air@latest

WORKDIR /go/src/app

COPY . .

RUN go mod tidy