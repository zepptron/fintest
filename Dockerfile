FROM golang:1.9-alpine
WORKDIR /go/src/
COPY code/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/main /usr/bin/main
COPY code/config.file /etc/financialcom/
CMD ["/usr/bin/main"]
