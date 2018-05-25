#!/bin/bash

if [[ $COUNTRY = "" ]]; then
    echo "COUNTRY is required"
    echo "example: COUNTRY=jp ./update.sh"
    exit 1
fi;

git reset

echo "try to update COUNTRY $COUNTRY"

sleep 2

if [[ $COUNTRY = "jp" ]]; then
    ./japanpost/gen.sh
else
    echo "COUNTRY $COUNTRY not supported"
    exit 1
fi;

if [[ $(git add $COUNTRY/ -v | wc -l) -gt 0 ]]; then
    DATE="$(date +"%Y-%m-%d")"
    git checkout -b "update-$COUNTRY-zipcode-on-$DATE"
    git commit -m "zipcode: update $COUNTRY on $DATE"
else
    echo "nothing update in $COUNTRY/ folder"
fi;
