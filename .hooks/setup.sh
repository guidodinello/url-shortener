#!/bin/bash

# Move to root directory
cd "$(dirname "$0")/.." || exit

# Make files executable
chmod +x .hooks/pre-commit
chmod +x .hooks/pre-push

# Move files to .git/hooks
REL_PATH_TARGET_TO_SRC="../../.hooks"
ln -s ${REL_PATH_TARGET_TO_SRC}/pre-commit .git/hooks/pre-commit
ln -s ${REL_PATH_TARGET_TO_SRC}/pre-push .git/hooks/pre-push
