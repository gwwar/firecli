# syntax=docker/dockerfile:1
FROM golang:1.13
# Create /app directory in image, and set it as the working directory (.)
WORKDIR /app

COPY ./ /app
# Download 3rd party dependencies
RUN go mod download

ENTRYPOINT go run main.go

#FROM prom/prometheus:v2.17.0

