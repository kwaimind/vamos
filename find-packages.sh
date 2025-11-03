#!/bin/bash

# Find all package.json files and show tree structure
if command -v tree &> /dev/null; then
    # Use tree if available
    output=$(tree -P "package.json" --prune -f -I "node_modules")
else
    # Fallback to find with formatting
    output=$(find . -name "package.json" -type f -not -path "*/node_modules/*" | sort)
fi

echo "$output" | pbcopy
echo "$output"
echo ""
echo "✓ Copied to clipboard"
