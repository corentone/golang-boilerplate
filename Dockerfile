FROM golang:1.9-alpine as builder
ADD . /go/src/github.com/corentone/golang-boilerplate/
RUN go install github.com/corentone/golang-boilerplate/cmd/...

FROM alpine:latest
COPY --from=builder /go/bin/* /boilerplate/
WORKDIR /boilerplate
CMD ["/boilerplate/boilerplatebin"]