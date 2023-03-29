#!/bin/sh
tcpdump -A -i eth0 src port 53 -w "attack.pcap" -W 1 -G 10

BITRATE=$(capinfos -i -T -b "attack.pcap" | tail +2 | awk '{print $2}')
echo Average bitrate: $BITRATE bits/sec

