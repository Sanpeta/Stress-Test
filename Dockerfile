#build stage
FROM golang:1.22.3-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

#run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

#permissões do executável
RUN chmod +x /app/main

#run command
ENTRYPOINT ["/app/main"]
