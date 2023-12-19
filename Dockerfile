FROM golang:latest as builder
RUN apt-get update && apt-get install -y bash
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/fart/main.go

FROM ubuntu
LABEL maintainer="Artur Zagirov"
COPY --from=builder main /main
EXPOSE 11111
ENTRYPOINT ["/main"]
