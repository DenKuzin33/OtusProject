FROM golang:latest as build

WORKDIR /server

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM scratch

EXPOSE 55011

COPY --from=build server .

ENTRYPOINT ["/server"]