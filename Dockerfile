
# First Stage: Build the Go App
FROM golang:1.23.4 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy 
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

# Second Stage: Add Ca Certificates
FROM alpine:latest as ca-certificates
RUN apk --update add ca-certificates

# Third Stage: Execute the Go App
FROM scratch
COPY --from=ca-certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /app
COPY --from=builder /app/main ./
EXPOSE 8080
CMD [ "./main" ]