FROM golang:1-bullseye AS builder

WORKDIR /app

COPY . .

RUN go build .

FROM debian:bullseye AS runtime

WORKDIR /app

COPY --from=builder /app/train-api .
COPY configurations .

RUN apt-get update -y && apt-get install -y curl gzip

ENV HTTP_HOSTNAME="0.0.0.0"
ENV HTTP_PORT="5000"

EXPOSE 5000

CMD ["/app/train-api"]