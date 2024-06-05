FROM golang:1.22.4-alpine3.19 as builder
WORKDIR /gotemplate
RUN apk update && apk upgrade --available && sync && apk add --no-cache --virtual .build-deps
COPY . .
RUN go build -ldflags="-w -s" .
FROM alpine:3.20.0
RUN apk update && apk upgrade --available && sync
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
