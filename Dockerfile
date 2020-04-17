FROM golang:1.11-alpine
WORKDIR /go/src/
COPY code/main.go .
RUN apk add --no-cache git && \
	go get github.com/prometheus/client_golang/prometheus && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM scratch
COPY --from=0 /go/src/main /main
COPY code/config.file /config.file
WORKDIR /
CMD ["/main"]
