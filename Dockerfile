FROM registrysecaas.azurecr.io/secaas/golang-dev:1.12-latest

WORKDIR /app
USER root
# ADD . app

COPY . /go/
# RUN ls -la
RUN pwd
RUN cd /go/src/
RUN go mod init
RUN go mod tidy
RUN go build /go/src/cmd/main.go
# RUN ls -la /go
# RUN ls -la /go/src
RUN ls -la /root
# RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app main.go

# FROM alpine:latest
# # mailcap adds mime detection and ca-certificates help with TLS (basic stuff)
# RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
# USER app
# WORKDIR /app

#COPY --from=builder /app/app .

# ENTRYPOINT ["go", "run", "/go/src/main.go"]
ENTRYPOINT ["/app/src/cmd/main"]
# CMD ["tail -f /dev/null"]
