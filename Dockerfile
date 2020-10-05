FROM golang:1.14-buster

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o invoices-api ./cmd/main.go

CMD ["./invoices-api"]