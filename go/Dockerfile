FROM alpine:latest

RUN apk update && apk add --no-cache tcpdump

ENTRYPOINT ["tcpdump"]
CMD ["-i", "eth0"]

