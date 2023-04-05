FROM golang:1.20.3-alpine3.17 as builder
WORKDIR /gotemplate
RUN apk update && apk upgrade --available && sync
COPY . .
RUN go build -ldflags="-w -s" .
RUN rm -rf *.go && rm -rf go.*
FROM alpine:3.17.3
RUN apk update && apk upgrade --available && sync
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
