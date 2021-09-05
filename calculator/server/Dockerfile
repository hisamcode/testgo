# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY ../../go.mod ./
COPY ../../go.sum ./
RUN go mod download

COPY server.go ./server/
COPY calculatorpb/*.go ./calculatorpb/

RUN go build -o /server-calculator

EXPOSE 8080

CMD [ "/server-calculator" ]