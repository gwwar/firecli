# syntax=docker/dockerfile:1
FROM golang:1.17
# Create /app directory in image, and set it as the working directory (.)
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
# Download 3rd party dependencies
RUN go mod download

COPY *.go ./
RUN go build -o /demoapp

EXPOSE 2112

CMD ["/demoapp"]
