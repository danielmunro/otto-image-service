FROM golang:1.13
WORKDIR /go/src/github.com/danielmunro/otto-image-service
COPY . .
RUN go get -d -v ./...
RUN go build
EXPOSE 8082
ENTRYPOINT ["./otto-image-service"]
