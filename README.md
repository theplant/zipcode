# Find Address By Zipcode

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
