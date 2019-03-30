FROM golang:alpine as builder

ADD . /go/src/github.com/thiagotrennepohl/fortune-scrapper
ENV GO111MODULE on
WORKDIR /go/src/github.com/thiagotrennepohl/fortune-scrapper
RUN apk add --update git && go mod download && CGO_ENABLED=0 go build -a -installsuffix main.go -o main

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/thiagotrennepohl/fortune-scrapper/main /app/main
ENTRYPOINT ./main