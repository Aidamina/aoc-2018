FROM golang:1.11.2-alpine AS builder
RUN apk update && apk add --no-cache git ca-certificates
WORKDIR /go/src/
COPY src/ /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app .
CMD ["./app"] 