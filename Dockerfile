FROM golang:1.24.2-alpine3.21 AS builder
WORKDIR /gotemplate
RUN apk add --no-cache ca-certificates
COPY . .
RUN go build -trimpath -ldflags="-w -s" .
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
