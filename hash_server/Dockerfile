FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app/
COPY . .
RUN go get ./...
RUN go build -o ./bin/GoHashServer ./cmd/hash_server/hashServer.go

FROM alpine:3.9
WORKDIR /go/bin
COPY --from=build /go/src/app/bin/GoHashServer /go/bin/GoHashServer
ENTRYPOINT /go/bin/GoHashServer
