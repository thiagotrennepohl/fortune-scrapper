FROM golang:alpine as builder

ADD . /go/src/github.com/thiagotrennepohl/fortune-scrapper
ENV GO111MODULE on
WORKDIR /go/src/github.com/thiagotrennepohl/fortune-scrapper
RUN apk add --update git ca-certificates && go mod download && CGO_ENABLED=0 go build -a -installsuffix main.go -o main

# final stage
FROM alpine
WORKDIR /app
RUN apk add --update ca-certificates
COPY --from=builder /go/src/github.com/thiagotrennepohl/fortune-scrapper/main /app/main
ENTRYPOINT ./main