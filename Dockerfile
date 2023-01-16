# syntax=docker/dockerfile:1

FROM golang:1.19.5-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY *.yaml ./

RUN go build -o /spec-api

EXPOSE 8080

CMD [ "/spec-api" ]