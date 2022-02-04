FROM golang:1.17

WORKDIR /app

COPY . .

RUN go get -u golang.org/x/lint/golint
RUN golint -set_exit_status .

RUN go mod download
RUN go test

RUN go build -o ping
CMD ["/app/ping"]

