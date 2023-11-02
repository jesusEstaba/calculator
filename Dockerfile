FROM golang:1.19

WORKDIR /go/src/github.com/jesusEstaba/calculator

COPY . .
RUN ls


RUN go build -o main ./cmd/api
EXPOSE 8080
CMD ["./main"]