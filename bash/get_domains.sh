#!/bin/bash

# Read input file
URLS=$(tail -n +2 input.csv | awk -F, '{print $1}')

OUTPUT=""

# Iterate over URLs in the csv
for URL in $URLS
do
    # Extract domain
    domain=$(echo $URL | awk -F[/:] '{print $4}')

    # Split the domain into its levels
    IFS='.' read -ra levels <<< "$domain"

    # Print each level of the domain and all lower levels
    for ((i=0; i<${#levels[@]}-1; i++)); do
        echo $(IFS=.; echo "${levels[*]:i}")
    done
done
