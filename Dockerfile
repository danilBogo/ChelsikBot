FROM golang:latest

# Установка пакета для работы с переменными окружения
RUN apt-get update && apt-get install -y gettext-base

WORKDIR /go/src/app

COPY . .

ARG TELEGRAM_BOT_TOKEN_ARG
ARG PINGS_ARG

# Проверяем наличие файла .env и создаем его, если он отсутствует
RUN test -f .env || echo "TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN_ARG}" > .env && \
    echo "PINGS=${PINGS_ARG}" >> .env

WORKDIR /go/src/app/cmd

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]