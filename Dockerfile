FROM alpine:latest

WORKDIR /app
COPY src/cmd/main .

ENTRYPOINT ["/app/main"]
