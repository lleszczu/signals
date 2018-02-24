FROM golang:1.9.3
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/os-signals
COPY . /go/src/os-signals
RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /os-signals ./cmd/main.go

FROM alpine
COPY --from=0 /os-signals .
ENTRYPOINT /os-signals
