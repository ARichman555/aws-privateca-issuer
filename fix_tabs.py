#!/usr/bin/env python3

import re

# Read the Makefile
with open('Makefile', 'r') as f:
    content = f.read()

# Replace leading 8 spaces with tabs for recipe lines
# This pattern matches lines that start with exactly 8 spaces
lines = content.split('\n')
fixed_lines = []

for line in lines:
    # Check if line starts with exactly 8 spaces followed by non-whitespace
    if re.match(r'^        [^ ]', line):
        # Replace the 8 leading spaces with a tab
        fixed_line = '\t' + line[8:]
        fixed_lines.append(fixed_line)
    else:
        fixed_lines.append(line)

# Write the fixed content back
with open('Makefile', 'w') as f:
    f.write('\n'.join(fixed_lines))

print("Fixed Makefile indentation - replaced leading 8 spaces with tabs")