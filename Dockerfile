FROM golang:1.22.3-alpine3.19 as builder
WORKDIR /gotemplate
RUN apk update && apk upgrade --available && sync && apk add --no-cache --virtual .build-deps
COPY . .
RUN go build -ldflags="-w -s" .
FROM alpine:3.19.1
RUN apk update && apk upgrade --available && sync
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
