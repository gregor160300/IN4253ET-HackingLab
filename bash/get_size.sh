DOMAINS=$(tail -n +2 domains_and_nameservers.csv | awk -F, '{print $1}')
NAME_LIST=$(tail -n +2 domains_and_nameservers.csv | awk -F, '{print $2}')

DOMAINS=($DOMAINS)
COUNT=0
for NAMES in $NAME_LIST
do
    DOMAIN=${DOMAINS[COUNT]}
    COUNT=$((COUNT+1))
    NAMES=$(echo $NAMES | tr ";" " ")
    # Request size is 60 + length domain name
    for NAME in $NAMES
    do
        REQUEST=$(echo ${#NAME})
        REQUEST=$((REQUEST+60))
        RESPONSE=$(host -U -v -t any -4 $DOMAIN $NAME | tail -n 1 | awk '{print $2}')
        echo "$DOMAIN,$NAME,$REQUEST,$RESPONSE"
    done
done

