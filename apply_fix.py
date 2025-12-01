#!/usr/bin/env python3
import re
import os

os.chdir('/workspace')

# Read the Makefile
with open('Makefile', 'r') as f:
    content = f.read()

print("Original content sample (line 70):")
lines = content.split('\n')
if len(lines) > 69:
    print(f"Line 70: {repr(lines[69])}")

# Create backup
with open('Makefile.backup', 'w') as f:
    f.write(content)

print("Created backup")

# Replace all lines that start with exactly 8 spaces with TAB
# Use regex to match start of line + 8 spaces + non-space character
fixed_content = re.sub(r'^        ([^\s])', r'\t\1', content, flags=re.MULTILINE)

# Count changes
original_lines = content.split('\n')
fixed_lines = fixed_content.split('\n')
changes = sum(1 for i, (orig, fixed) in enumerate(zip(original_lines, fixed_lines)) if orig != fixed)

print(f"Fixed {changes} lines")

# Show sample of fixed line
if len(fixed_lines) > 69:
    print(f"Fixed Line 70: {repr(fixed_lines[69])}")

# Write the fixed content
with open('Makefile', 'w') as f:
    f.write(fixed_content)

print("Applied TAB indentation fix")

# Now fix the lint target
with open('Makefile', 'r') as f:
    content = f.read()

# Fix lint target
old_lint_pattern = r'lint:\s*\n\s*echo "Linter is deprecated with go1\.18!"'
new_lint = 'lint: golangci-lint\n\t$(GOLANGCILINT) run --timeout 10m'

if re.search(old_lint_pattern, content):
    content = re.sub(old_lint_pattern, new_lint, content)
    print("Updated lint target to use golangci-lint")
else:
    print("Lint target pattern not found")

# Update golangci-lint version
old_version = 'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2'
new_version = 'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2'

if old_version in content:
    content = content.replace(old_version, new_version)
    print("Updated golangci-lint version to v1.55.2")
else:
    print("golangci-lint version not found or already updated")

# Write the final content
with open('Makefile', 'w') as f:
    f.write(content)

print("Makefile has been updated with proper TAB indentation and golangci-lint integration")
print("All fixes applied successfully!")