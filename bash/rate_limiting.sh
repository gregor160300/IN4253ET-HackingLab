#!/bin/bash
counter=0
while true
do
     counter=$((counter+1))
     # dig -4 +notcp +ignore +bufsize=4096 @nameserver domain_name ANY
     dig -4 +notcp +ignore +bufsize=4096 @ns1.rijksoverheidnl.nl. iplo.nl ANY

     echo "Queries: $counter"
done
