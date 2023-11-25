#!/bin/bash

# Move to root directory
cd "$(dirname "$0")/.." || exit

# Move files to .git/hooks
cp .hooks/pre-commit .git/hooks/pre-commit
cp .hooks/pre-push .git/hooks/pre-push

# Make files executable
chmod +x .git/hooks/pre-commit
chmod +x .git/hooks/pre-push
