FROM alpine:latest

WORKDIR /app
COPY src/main .

ENTRYPOINT ["/app/main"]
