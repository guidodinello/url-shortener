#!/bin/bash

# Move to root directory
cd "$(dirname "$0")/.." || exit

# Make files executable
chmod +x .hooks/pre-commit
chmod +x .hooks/pre-push

# Move files to .git/hooks
ln -s .hooks/pre-commit .git/hooks/pre-commit
ln -s .hooks/pre-push .git/hooks/pre-push
