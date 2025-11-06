#!/bin/bash

# This script forces the shell to recognize the new command and then runs it.

# Load the latest shell configuration
CONFIG_FILE=""
if [ -f "$HOME/.zshrc" ]; then
    CONFIG_FILE="$HOME/.zshrc"
elif [ -f "$HOME/.bashrc" ]; then
    CONFIG_FILE="$HOME/.bashrc"
fi

if [ -n "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
fi

# Explicitly add GOPATH/bin to the PATH for this script's execution context
export PATH=$PATH:$(go env GOPATH)/bin

# Execute the gitprofile command
exec gitprofile "$@"