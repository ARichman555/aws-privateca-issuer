#!/bin/bash

# Fix Makefile indentation by replacing 8 spaces with tabs for recipe lines
# This script identifies lines that should be recipe lines (commands under targets)
# and converts their leading 8-space indentation to tab characters

# Create a backup
cp Makefile Makefile.backup

# Use sed to replace leading 8 spaces with a tab character
# This pattern matches lines that start with exactly 8 spaces followed by non-whitespace
sed -i 's/^        \([^ ]\)/\t\1/' Makefile

echo "Fixed Makefile indentation - replaced leading 8 spaces with tabs"
echo "Backup saved as Makefile.backup"