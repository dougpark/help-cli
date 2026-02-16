#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Configuration
SOURCE_NAME="help.go"
APP_NAME="help"
OUT_DIR="./dist"
PROD_BIN="$HOME/bin"

# 1. Get Build Metadata
CURRENT_DATE=$(date +"%Y-%m-%d %H:%M:%S")

# Get the short 7-character Git hash (if in a git repo)
GIT_HASH=$(git rev-parse --short HEAD 2>/dev/null || echo "not-git")

# 2. Linker flags
# We add a second -X flag to inject the gitHash into your Go code
LDFLAGS="-s -w -X 'main.buildDate=$CURRENT_DATE' -X 'main.gitHash=$GIT_HASH'"

# Parse flags
DEPLOY_PROD=false
while getopts "p" opt; do
  case $opt in
    p) DEPLOY_PROD=true ;;
    *) echo "Usage: $0 [-p]"; exit 1 ;;
  esac
done

echo "Building version: $GIT_HASH ($CURRENT_DATE)"

# 3. Clean up
rm -rf $OUT_DIR
mkdir -p $OUT_DIR

# 4. Build for macOS (ARM64)
# Note: We include the OpenSSL paths for your M3 build
echo "Building macOS..."
CGO_CPPFLAGS="-I/opt/homebrew/opt/openssl@3/include" \
CGO_LDFLAGS="-L/opt/homebrew/opt/openssl@3/lib" \
go build -ldflags "$LDFLAGS" -o $OUT_DIR/$APP_NAME ./$SOURCE_NAME

# 5. Build for Linux Server (AMD64)
echo "Building Linux..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFLAGS" -o $OUT_DIR/$APP_NAME-linux ./$SOURCE_NAME

# 6. Build for Raspberry Pi (ARM64)
echo "Building Pi..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$LDFLAGS" -o $OUT_DIR/$APP_NAME-pi ./$SOURCE_NAME

echo "---------------------------------------"
echo "Build Successful! Binaries in $OUT_DIR"

if [ "$DEPLOY_PROD" = true ]; then
    mkdir -p "$PROD_BIN"
    cp "$OUT_DIR/$APP_NAME" "$PROD_BIN/"
    echo "Deployed to $PROD_BIN/$APP_NAME"
fi