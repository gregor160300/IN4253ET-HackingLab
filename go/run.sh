#!/bin/sh
# DO NOT USE A 1 in POS 3, it cannot be spoofed for some reason
TARGET_IP=192.168.1.10
go run main.go --target_ip=$TARGET_IP --duration=30

