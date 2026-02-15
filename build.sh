#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Configuration
SOURCE_NAME="help.go"
APP_NAME="help"
OUT_DIR="./dist"
PROD_BIN="$HOME/bin"

# Get current date for the binary
CURRENT_DATE=$(date +"%Y-%m-%d %H:%M:%S")
# Linker flags: -s -w reduces binary size, -X injects the date
LDFLAGS="-s -w -X 'main.buildDate=$CURRENT_DATE'"

# Parse flags
DEPLOY_PROD=false
while getopts "p" opt; do
  case $opt in
    p) DEPLOY_PROD=true ;;
    *) echo "Usage: $0 [-p]"; exit 1 ;;
  esac
done

echo "Building version dated: $CURRENT_DATE"

# 1. Clean up
rm -rf $OUT_DIR
mkdir -p $OUT_DIR

# 2. Build for macOS (ARM64)
echo "Building macOS..."
go build -ldflags "$LDFLAGS" -o $OUT_DIR/$APP_NAME ./$SOURCE_NAME

# 3. Build for Linux Server (AMD64)
echo "Building Linux..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFLAGS" -o $OUT_DIR/$APP_NAME-linux ./$SOURCE_NAME

# 4. Build for Raspberry Pi (ARM64)
echo "Building Pi..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$LDFLAGS" -o $OUT_DIR/$APP_NAME-pi ./$SOURCE_NAME

echo "---------------------------------------"
echo "Build Successful! Binaries in $OUT_DIR"

if [ "$DEPLOY_PROD" = true ]; then
    mkdir -p "$PROD_BIN"
    cp "$OUT_DIR/$APP_NAME" "$PROD_BIN/"
    echo "Deployed to $PROD_BIN/$APP_NAME"
fi