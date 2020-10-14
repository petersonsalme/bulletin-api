FROM golang:latest

RUN mkdir -p /go/src/bulletin-api

WORKDIR /go/src/bulletin-api

COPY . /go/src/bulletin-api

RUN go build 

RUN go install 

CMD bulletin-api

EXPOSE 8080