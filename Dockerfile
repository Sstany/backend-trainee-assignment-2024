FROM golang:1.21.0-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk update && \
    apk add build-base

RUN go mod download

COPY ./ ./

RUN go build -o banney /app/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/banney /app/banney

EXPOSE 8090

CMD [ "./banney" ]
