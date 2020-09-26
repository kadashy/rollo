FROM alpine:latest

WORKDIR /app
COPY main .

ENTRYPOINT ["/app/main"]
