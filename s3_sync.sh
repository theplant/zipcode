#!/bin/bash

echo $BUCKET

aws s3 sync zipcode/ s3://$BUCKET/zipcode/ --exclude "*" --include "*.json"

