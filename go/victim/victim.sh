#!/bin/sh
# tcpdump -A -i eth0 -c 10000 src port 53 -w attack.pcap
tcpdump -i eth0 src port 53 -c 1000 -vvv

BITRATE=$(capinfos -i -T -b attack.pcap | tail +2 | awk '{print $2}')
echo Average bitrate: $BITRATE bits/sec

