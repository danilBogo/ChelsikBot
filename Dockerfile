FROM golang:latest

WORKDIR /go/src/app

COPY . .

WORKDIR /go/src/app/cmd

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]