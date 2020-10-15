FROM golang:alpine3.12

ENV BULLETIN_API=/bulletin-api
ENV MIGRATIONS_PATH=/migrations

ENV GOOS=linux 
ENV GOARCH=386 

RUN mkdir -p ${MIGRATIONS_PATH}
COPY ./migrations ${MIGRATIONS_PATH}

RUN mkdir -p ${BULLETIN_API}
ADD go.mod ${BULLETIN_API}
ADD go.sum ${BULLETIN_API}
COPY ./api ${BULLETIN_API}/api
COPY ./pkg ${BULLETIN_API}/pkg

WORKDIR ${BULLETIN_API}

RUN go mod download
RUN go build -o ./bin/bulletin-api api/main.go

CMD ./bin/bulletin-api

EXPOSE 8080