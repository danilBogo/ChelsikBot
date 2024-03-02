FROM golang:latest

WORKDIR /go/src/app

COPY . .

WORKDIR /go/src/app/cmd

ARG TELEGRAM_BOT_TOKEN_ARG
ARG PINGS_ARG

ENV TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN_ARG
ENV PINGS=$PINGS_ARG

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]