FROM golang:1.17 as base

WORKDIR /app

COPY . .

FROM base as linter
RUN go get -u golang.org/x/lint/golint
RUN golint -set_exit_status .

FROM base as test
RUN go mod download
RUN go test

FROM test as builder
RUN CGO_ENABLED=0 go build -o ping

FROM alpine as final
COPY --from=builder /app/ping /ping

CMD ["/ping"]

