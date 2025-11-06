#!/bin/bash

# Exit on any error
set -e

echo "Installing gitprofile..."

# 1. Run go install
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi
go install

# 2. Detect shell and set config file path
SHELL_TYPE=$(basename "$SHELL")
CONFIG_FILE=""
if [ "$SHELL_TYPE" = "zsh" ]; then
    CONFIG_FILE="$HOME/.zshrc"
elif [ "$SHELL_TYPE" = "bash" ]; then
    CONFIG_FILE="$HOME/.bashrc"
else
    echo "Unsupported shell: $SHELL_TYPE. Please add the Go bin directory to your PATH manually."
    exit 1
fi

# 3. Get Go path and ensure it's not empty
GOPATH_BIN=$(go env GOPATH)/bin
if [ -z "$GOPATH_BIN" ]; then
    echo "Could not determine Go path. Please ensure Go is configured correctly."
    exit 1
fi

# 4. Check if path is already in the config file
if ! grep -q "export PATH=\$PATH:$GOPATH_BIN" "$CONFIG_FILE"; then
    echo "Adding Go binary path to $CONFIG_FILE"
    # Append the export command to the config file
    echo -e "\\n# Add Go binaries to path for gitprofile" >> "$CONFIG_FILE"
    echo "export PATH=\$PATH:$GOPATH_BIN" >> "$CONFIG_FILE"
    echo "Successfully updated shell configuration."
else
    echo "Go binary path is already in your shell configuration."
fi

# 5. Check for existing profiles
PROFILE_CONFIG_PATH="$HOME/.config/gitprofile/profiles.json"
if [ -f "$PROFILE_CONFIG_PATH" ]; then
    echo -e "\\nFound existing profiles! Listing them for you:"
    "$GOPATH_BIN/gitprofile" list
else
    echo -e "\\nNo existing profiles found. Get started by running: gitprofile add"
fi

# 6. Final instruction
echo -e "\\n✅ Installation complete!"
echo "Please restart your terminal or run the following command to use 'gitprofile':"
echo "source $CONFIG_FILE"