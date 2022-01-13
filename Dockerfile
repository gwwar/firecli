# syntax=docker/dockerfile:1
FROM golang:1.17
# Create /app directory in image, and set it as the working directory (.)
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
# Download 3rd party dependencies
RUN go mod download

COPY main.go ./
COPY cmd ./cmd
COPY assets ./assets

RUN go install main.go

ENTRYPOINT ["main"]
