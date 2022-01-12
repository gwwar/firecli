# syntax=docker/dockerfile:1
FROM golang:1.17
# Create /app directory in image, and set it as the working directory (.)
WORKDIR /app

COPY ./ /app
# Download 3rd party dependencies
RUN go mod download
RUN go install main.go

ENTRYPOINT ["main"]
