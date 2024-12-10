FROM alpine:latest

WORKDIR /app

RUN apk update && apk add --no-cache \
    tor \
    curl \
    bash \
    libevent

COPY dist/server .

COPY torrc /etc/tor/torrc

EXPOSE 8080 9050

CMD tor & ./server
