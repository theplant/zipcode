#!/bin/bash

if [[ $BUCKET = "" || $COUNTRY = "" ]]; then
    echo "both of BUCKET and COUNTRY are required"
    echo "example: BUCKET=zipcode.domain.com COUNTRY=jp ./s3_sync.sh"
    exit 1
fi;

echo "BUCKET $BUCKET"
echo "COUNTRY $COUNTRY"

tar -xzf ./$COUNTRY.tar.gz && aws s3 sync $COUNTRY/ s3://$BUCKET/$COUNTRY/ --exclude "*" --include "*.json"
