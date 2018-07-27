FROM golang:1.9-alpine
WORKDIR /go/src/
COPY code/main.go .
RUN apk add --no-cache git && \
	go get github.com/prometheus/client_golang/prometheus && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:latest
COPY --from=0 /go/src/main /usr/bin/main
COPY code/config.file /etc/financialcom/
WORKDIR /etc/financialcom/
CMD ["/usr/bin/main"]
