FROM golang:latest

COPY . /app
WORKDIR /app

RUN apt update
RUN apt install -y libpcap* net-tools iputils-ping tcpdump
RUN go install

ENTRYPOINT ["sh", "./run.sh"]

