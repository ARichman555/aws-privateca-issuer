#!/usr/bin/env python3
import os
import subprocess

os.chdir('/workspace')

# Read the current Makefile
with open('Makefile', 'r') as f:
    content = f.read()

# Create backup
with open('Makefile.backup', 'w') as f:
    f.write(content)

print("Created backup: Makefile.backup")

# Fix the indentation: replace 8 spaces at start of line with single TAB
lines = content.split('\n')
fixed_lines = []

for i, line in enumerate(lines):
    # Check if line starts with exactly 8 spaces (recipe line)
    if line.startswith('        ') and not line.startswith('         '):  # 8 spaces, not 9+
        # Replace the 8 spaces with a single TAB
        fixed_line = '\t' + line[8:]
        fixed_lines.append(fixed_line)
        print(f"Fixed line {i+1}: {repr(line[:20])} -> {repr(fixed_line[:20])}")
    else:
        fixed_lines.append(line)

# Write the fixed content
fixed_content = '\n'.join(fixed_lines)

# Now also fix the lint target and golangci-lint version as intended in the original PR
fixed_content = fixed_content.replace(
    'lint:\n\techo "Linter is deprecated with go1.18!"',
    'lint: golangci-lint\n\t$(GOLANGCILINT) run --timeout 10m'
)

# Update golangci-lint version to v1.55.2 as intended in the original PR
fixed_content = fixed_content.replace(
    'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2',
    'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2'
)

with open('Makefile', 'w') as f:
    f.write(fixed_content)

print(f"\nFixed {len([l for l in lines if l.startswith('        ') and not l.startswith('         ')])} recipe lines")
print("Updated lint target to use golangci-lint")
print("Updated golangci-lint version to v1.55.2")
print("Makefile has been updated with proper TAB indentation and golangci-lint integration")

# Test the syntax
print("\nTesting Makefile syntax...")
try:
    result = subprocess.run(['make', '--dry-run', 'test'], capture_output=True, text=True, timeout=10)
    if result.returncode == 0:
        print("✓ Makefile syntax is now correct!")
    else:
        print(f"✗ Makefile syntax error: {result.stderr}")
except Exception as e:
    print(f"Error testing make: {e}")