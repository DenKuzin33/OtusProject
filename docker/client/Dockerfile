FROM golang:latest as build

WORKDIR /client

COPY go.mod ./
COPY *.go ./

RUN go build -o /client

EXPOSE 5555

ENTRYPOINT [ "/client" ]