#!/usr/bin/env python3
import re

# Read the Makefile
with open('/workspace/Makefile', 'r') as f:
    content = f.read()

# Create backup
with open('/workspace/Makefile.backup', 'w') as f:
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

# Also fix the lint target and golangci-lint version
fixed_content = fixed_content.replace(
    'lint:\n\techo "Linter is deprecated with go1.18!"',
    'lint: golangci-lint\n\t$(GOLANGCILINT) run --timeout 10m'
)

fixed_content = fixed_content.replace(
    'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2',
    'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2'
)

# Write the fixed content
with open('/workspace/Makefile', 'w') as f:
    f.write(fixed_content)

print("Makefile fixed with TAB indentation and golangci-lint integration")

# Test syntax
import subprocess
try:
    result = subprocess.run(['make', '--dry-run', 'test'], 
                          capture_output=True, text=True, timeout=10, cwd='/workspace')
    if result.returncode == 0:
        print("✓ Makefile syntax is correct!")
    else:
        print(f"✗ Error: {result.stderr}")
except Exception as e:
    print(f"Error testing: {e}")