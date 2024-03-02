# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY *.go ./

RUN go build ./cmd/app -o /hackiwna-docker

EXPOSE 8080

CMD ["/hackiwna-docker"]
