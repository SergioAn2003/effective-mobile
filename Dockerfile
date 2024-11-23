FROM golang:1.23.3-alpine3.20 AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /app/main ./cmd/main.go

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8000
ENTRYPOINT [ "./main" ]
