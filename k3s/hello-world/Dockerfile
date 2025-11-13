FROM golang:1.25.4-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /hello-world

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /hello-world .
EXPOSE 8080
CMD ["./hello-world"]
