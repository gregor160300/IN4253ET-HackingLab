DOMAINS=$(tail -n +2 domains_and_nameservers.csv | awk -F, '{print $1}')
NAME_LIST=$(tail -n +2 domains_and_nameservers.csv | awk -F, '{print $2}')

DOMAINS=($DOMAINS)
COUNT=0
echo "domain,name,request,response,truncated"
for NAMES in $NAME_LIST
do
    DOMAIN=${DOMAINS[COUNT]}
    COUNT=$((COUNT+1))
    NAMES=$(echo $NAMES | tr ";" " ")
    for NAME in $NAMES
    do
        REQUEST=$(echo ${#DOMAIN})
        REQUEST=$((REQUEST+83))
        OUTPUT=$(dig -4 +notcp +ignore +bufsize=4096 @$NAME $DOMAIN ANY)
        TC_FLAG=$(echo "$OUTPUT" | grep -o "\btc\b" )
        RESPONSE=$(echo "$OUTPUT" | tail -n 2 | grep -E -o "\brcvd: [0-9]+\b" | awk '{print $2}')
        if [[ "$TC_FLAG" == "tc" ]]
        then
            echo "$DOMAIN,$NAME,$REQUEST,$RESPONSE,yes"
        else
            echo "$DOMAIN,$NAME,$REQUEST,$RESPONSE,no"
        fi
        sleep 1
    done
done
