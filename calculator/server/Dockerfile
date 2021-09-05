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

FROM <image_name>:<image_version>
WORKDIR /app
COPY . .
RUN echo "Installing or doing stuff"
RUN <my_command>
EXPOSE <app_port>
CMD [ "<command>", "<argument_1>", "<argument_2>" ]