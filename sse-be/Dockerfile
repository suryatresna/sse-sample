FROM golang:1.14

WORKDIR /go/src/github.com/suryatresna/sse-sample/sse-be

COPY . .

RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM scratch
COPY --from=0 /go/src/github.com/suryatresna/sse-sample/sse-be/main /main

CMD ["/main"]