# Read input file
DOMAINS=$(cat $1)

# Print csv headers to stdout
echo "domain,nameserver,ip,request,response,tc"

# Iterate over all URLs in the input file
for DOMAIN in $DOMAINS
do
    # Query nameservers for domain
    NAMES=$(dig +short NS $DOMAIN)
    for NAME in $NAMES
    do
        # Calculate request size
        REQUEST=$(echo ${#DOMAIN})
        REQUEST=$((REQUEST+71))
        # Perform dig command
        OUTPUT=$(dig -4 +notcp +ignore +bufsize=4096 @$NAME $DOMAIN ANY)
        # Parse output from dig command
        TC_FLAG=$(echo "$OUTPUT" | grep -o "\btc\b" )
        RESPONSE=$(echo "$OUTPUT" | tail -n 2 | grep -E -o "\brcvd: [0-9]+\b" | awk '{print $2}')
        NAME_IP=$(echo "$OUTPUT" | grep -E -o "SERVER: ([0-9]{1,3}[\.]){3}[0-9]{1,3}" | awk '{print $2}')
        # Print parsed output to stdout
        echo "$DOMAIN,$NAME,$NAME_IP,$REQUEST,$RESPONSE,$TC_FLAG"
    done
    # Small sleep in the script to not overload our or another network
    sleep 0.1
done
