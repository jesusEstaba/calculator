FROM golang:1.19

WORKDIR /go/src/github.com/jesusEstaba/calculator

COPY . .

RUN go build -o main ./cmd/seed

CMD ["./main"]