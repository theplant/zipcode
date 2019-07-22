#!/bin/bash

## Init work dir
DIR=$(pwd)
WORK_DIR=${DIR%/japanpost}

## Init tmp dir
TMP_DIR="$WORK_DIR/tmp"
rm -rf "$TMP_DIR"
mkdir -p "$TMP_DIR"

## Download ken_all_rome.zip
DOWNLOAD_FILE="ken_all_rome.zip"
curl -fsSL -o "$TMP_DIR/$DOWNLOAD_FILE" "http://www.post.japanpost.jp/zipcode/dl/roman/$DOWNLOAD_FILE"

## unzip ken_all_rome.zip
unzip "$TMP_DIR/$DOWNLOAD_FILE" -d "$TMP_DIR"

SOURCE_FILE="$TMP_DIR/KEN_ALL_ROME.CSV"
TARGET_DIR="$WORK_DIR/jp"

## Init target dir
rm -rf "$TARGET_DIR" && tar -xzvf $WORK_DIR/jp.tar.gz -C $WORK_DIR

$WORK_DIR/japanpost/processor --source-file=$SOURCE_FILE --target-dir=$TARGET_DIR --verbose=true

tar -czf $WORK_DIR/jp.tar.gz $WORK_DIR/jp
