# Find Address By Zipcode

[![Build Status](https://semaphoreci.com/api/v1/theplant/zipcode/branches/master/badge.svg)](https://semaphoreci.com/theplant/zipcode)

## Use our service

```
curl -X GET http://zipcode.theplant-dev.com/jp/7860056.json

{"prefecture":"高知県","city":"高岡郡　四万十町","town":"志和"}
```

## Update `COUNTRY/*.json` files

```
COUNTRY=jp ./update.sh
```

## Sync `COUNTRY/*.json` files to S3

```
aws configure

BUCKET=zipcode.domain.com COUNTRY=jp ./s3_sync.sh
```

## Country code

https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements
