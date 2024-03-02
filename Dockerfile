FROM golang:latest

RUN apt-get update && apt-get install -y gettext-base

WORKDIR /go/src/app

COPY . .

RUN test -f .env || echo "TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}" > .env && \
    echo "PINGS=${PINGS}" >> .env

WORKDIR /go/src/app/cmd

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
