FROM golang:1.20 AS builder
WORKDIR /app

RUN go mod init github.com/Teroaz/projet-devops
COPY . .
RUN go mod tidy

RUN go build -o app ./webapi/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
