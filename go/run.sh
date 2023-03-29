#!/bin/sh
GATEWAY_MAC=$(arp | grep _gateway | awk '{print $3}')
TARGET_IP=192.168.178.70

# Need sudo for accessing network interface
sudo go run main.go --target_ip=$TARGET_IP --gateway=$GATEWAY_MAC --duration=3

