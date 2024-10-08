# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN echo "Installing or doing stuff"
RUN go mod download

COPY calculator ./calculator

RUN cd calculator/server/ && go build

EXPOSE 9080

CMD [ "/app/calculator/server/server" ]