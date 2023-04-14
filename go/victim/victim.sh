#!/bin/sh
tcpdump -A -i eth0 -c 100000 dst port 42000 -w attack.pcap

BITRATE=$(capinfos -i -T -b attack.pcap | tail +2 | awk '{print $2}')
echo Average bitrate: $BITRATE bits/sec

