FROM golang:1.25.6-alpine3.22 AS builder
WORKDIR /gotemplate
RUN apk add --no-cache git ca-certificates
COPY . .
RUN go build -trimpath -ldflags="-w -s" .
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
