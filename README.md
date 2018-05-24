# Find Japan Address By Zipcode

## Update `zipcode/*.json` files

```
./update.sh
```

## Sync `zipcode/*.json` files to S3

```
aws configure

BUCKET=zipcode ./s3_sync.sh
```
