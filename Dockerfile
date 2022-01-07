# syntax=docker/dockerfile:1
FROM golang:1.13
# Create /app directory in image, and set it as the working directory (.)
WORKDIR /app
# Download 3rd party dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
# build binary to /test
RUN go build -o /test

CMD [ "/test" ]

#FROM prom/prometheus:v2.17.0

