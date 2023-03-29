FROM alpine:latest

RUN apk update && apk add --no-cache tcpdump wireshark

COPY . /

ENTRYPOINT ["sh", "/victim.sh"]

