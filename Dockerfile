FROM golang:alpine as builder

RUN adduser -D -g '' appuser

COPY . $GOPATH/src/github.com/ndwhtlssthr/httpbin-go/
WORKDIR $GOPATH/src/github.com/ndwhtlssthr/httpbin-go/cmd/httpbin-go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/httpbin-go

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/httpbin-go /go/bin/httpbin-go
USER appuser

ENTRYPOINT ["/go/bin/httpbin-go"]
