FROM golang:1.21.5 AS builder
WORKDIR /app

RUN go mod init github.com/Teroaz/projet-devops
COPY . .
RUN go mod tidy

RUN go build -o app ./webapi/main.go
RUN "ls -la"

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
RUN "ls -la"
EXPOSE 8080
CMD ["./app"]
