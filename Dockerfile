FROM golang:1.15-alpine3.12 AS payment-builder

WORKDIR /app
ADD . /app

RUN go build -o payment-challenge ./cmd/order-gokomodo

FROM alpine:3.12 as order-gokomodo-app

COPY --from=cas-builder /app/order-gokomodo /app/

CMD [ "/app/order-gokomodo" ]
