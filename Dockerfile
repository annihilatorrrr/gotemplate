FROM golang:1.24.5-alpine3.22 AS builder
WORKDIR /gotemplate
RUN apk add --no-cache ca-certificates
COPY . .
RUN go build -trimpath -ldflags="-w -s" .
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /gotemplate/gotemplate /gotemplate
ENTRYPOINT ["/gotemplate"]
