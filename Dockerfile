FROM golang:1.18-alpine as builder

WORKDIR /app
COPY . .

RUN go build -o labpi

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/labpi /usr/local/bin/labpi

ENTRYPOINT ["labpi"]
CMD ["--help"]
