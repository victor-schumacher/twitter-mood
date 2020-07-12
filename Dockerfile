FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

EXPOSE $PORT

COPY . .

