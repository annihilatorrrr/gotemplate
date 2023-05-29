FROM golang:1.20.4-alpine3.18 as builder
WORKDIR /gotemplate
RUN apk update && apk upgrade --available && sync && apk add --no-cache --virtual .build-deps upx
COPY . .
RUN go build -ldflags="-w -s" .
RUN rm -rf *.go go.* && upx /gotemplate/gotemplate && apk --purge del .build-deps
FROM alpine:3.18.0
RUN apk update && apk upgrade --available && sync
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
