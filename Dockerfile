FROM golang:1.12 


# build directories
RUN mkdir -p /go/src/github.com/go-hackernews/
ADD . /go/src/github.com/go-hackernews/
WORKDIR /go/src/github.com/go-hackernews
RUN go get -u github.com/golang/dep/...
RUN dep ensure -vendor-only
RUN go test ./... -v 
RUN go build -o hackernews ./cmd/main.go

ENTRYPOINT ["./hackernews"]
