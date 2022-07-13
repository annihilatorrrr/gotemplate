FROM golang:1.19rc2-alpine3.16 as builder
WORKDIR /gotemplate
RUN apk update && apk upgrade --available && sync
COPY . .
RUN go build -ldflags="-w -s" .
RUN rm -rf *.go && rm -rf go.*
FROM alpine:3.16.0
RUN apk update && apk upgrade --available && sync
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
