#!/bin/bash

## Init work dir
DIR=$(pwd)
WORK_DIR=${DIR%/japanpost}

## Init tmp dir
TMP_DIR="$WORK_DIR/tmp"
rm -rf "$TMP_DIR"
mkdir -p "$TMP_DIR"

## Download KEN_ALL_ROME.zip
DOWNLOAD_FILE="KEN_ALL_ROME.zip"
curl -fsSL -o "$TMP_DIR/$DOWNLOAD_FILE" "https://www.post.japanpost.jp/service/search/zipcode/download/roman/$DOWNLOAD_FILE"

## unzip KEN_ALL_ROME.zip
unzip "$TMP_DIR/$DOWNLOAD_FILE" -d "$TMP_DIR"

SOURCE_FILE="$TMP_DIR/KEN_ALL_ROME.CSV"
TARGET_DIR="$WORK_DIR/jp"

## Build the processor for the host platform when Go is available.
## The committed binary is architecture-specific (built for macOS/amd64),
## so rebuilding keeps this working on arm64 Macs and on Linux CI runners.
if command -v go >/dev/null 2>&1; then
    (cd "$WORK_DIR/japanpost" && go build -o processor main.go)
fi

## Init target dir
rm -rf "$TARGET_DIR" && tar -xzf "$WORK_DIR/jp.tar.gz" -C "$WORK_DIR"

"$WORK_DIR/japanpost/processor" --source-file="$SOURCE_FILE" --target-dir="$TARGET_DIR" --verbose=true

## Pack with a relative path so the archive stays rooted at jp/ (s3_sync.sh
## extracts it and runs `aws s3 sync jp/`, which needs a top-level jp/ entry).
## COPYFILE_DISABLE=1 stops macOS tar from adding AppleDouble (._*) members and
## xattr headers; otherwise GNU tar on Linux extracts the ._* files and they get
## synced to S3 as junk (doubling the upload).
COPYFILE_DISABLE=1 tar -C "$WORK_DIR" -czf "$WORK_DIR/jp.tar.gz" jp
