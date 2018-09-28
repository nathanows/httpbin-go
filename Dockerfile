FROM golang:alpine as builder

RUN adduser -D -g '' appuser

COPY . $GOPATH/src/github.com/ndwhtlssthr/arp/
WORKDIR $GOPATH/src/github.com/ndwhtlssthr/arp/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/arp

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/arp /go/bin/arp
USER appuser
ENTRYPOINT ["/go/bin/arp"]
