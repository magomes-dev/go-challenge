FROM golang:1.22.2

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy