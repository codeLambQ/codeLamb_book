FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o /bin/server ./cmd/server

FROM alpine:3.20
COPY --from=builder /bin/server /usr/local/bin/server
COPY configs /app/configs
EXPOSE 8080
ENTRYPOINT ["server"]

