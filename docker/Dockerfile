FROM golang:latest

RUN mkdir -p /bulletin-api/bin

COPY ./bin/bulletin-api /bulletin-api/bin

RUN mkdir -p /bulletin-api/migrations

COPY ./migrations /bulletin-api/migrations

CMD bulletin-api

EXPOSE 8080