FROM alpine:latest

RUN apk update && apk add --no-cache tcpdump wireshark

COPY . /victim
WORKDIR /victim

ENTRYPOINT ["sh", "./victim.sh"]

