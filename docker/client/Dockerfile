FROM golang:latest as build

WORKDIR /client

COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /client

FROM scratch

COPY --from=build client .

ENTRYPOINT ["/client"]