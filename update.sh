#!/bin/bash

git reset

./japanpost/gen.sh

if [$(git add zipcode/ -v | wc -l) -gt 0 ]; then
    DATE="$(date +"%Y-%m-%d")"
    git checkout -b "update-zipcode-on-$DATE"
    git commit -m "zipcode: update on $DATE"
else
    echo "nothing update in zipcode/ folder"
fi
