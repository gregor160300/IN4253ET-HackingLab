#!/bin/bash

# Check for the input file
if [ -z "$1" ]; then
    echo "Please provide the input CSV file name as the first argument"
    exit 1
fi

# Check for the output files
if [ -z "$2" ]; then
    echo "Please provide the output CSV file name as the second argument"
    exit 1
fi

if [ -z "$3" ]; then
    echo "Please provide the output file name for unique nameservers as the third argument"
    exit 1
fi

# Create the output file and write the header row
echo "URL,nameservers" > $2

# Declare an associative array to store unique nameservers
declare -A nameserver_map

# Loop over each line in the input file and lookup the nameservers
while IFS=, read -r URL; do
    # Remove "http://" or "https://"
    URL=${URL#http://}
    URL=${URL#https://}
    # Remove "www."
    URL=${URL#www.}
    # Lookup the nameservers
    nameservers=$(dig +short NS $URL | tr '\n' ';' | sed 's/,$//')
    # Store unique nameservers in the array
    IFS=';' read -ra ns_list <<< "$nameservers"
    for ns in "${ns_list[@]}"; do
        if [ -n "$ns" ]; then
            nameserver_map["$ns"]=1
        fi
    done
    # Write the URL and nameservers to the output file
    echo "$URL,$nameservers" >> $2
done < <(cut -d "," -f1,1 $1 | tail -n +2)

# Write the unique nameservers to a third output file
echo "nameserver" > $3
for ns in "${!nameserver_map[@]}"; do
    echo "$ns" >> $3
done
