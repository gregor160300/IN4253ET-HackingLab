#!/bin/sh
ping google.com -c 1
# Get gateway MAC
GATEWAY_MAC=$(arp | awk '/Address/ {getline; getline; print $3; exit}')
TARGET_IP=192.168.1.11

go run main.go --target_ip=$TARGET_IP --gateway=$GATEWAY_MAC --duration=10

