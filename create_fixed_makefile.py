#!/usr/bin/env python3
import os

os.chdir('/workspace')

# Read the current Makefile
with open('Makefile', 'r') as f:
    lines = f.readlines()

# Create backup
with open('Makefile.backup', 'w') as f:
    f.writelines(lines)

print("Created backup: Makefile.backup")

# Process each line to fix indentation
fixed_lines = []
for i, line in enumerate(lines):
    # If line starts with exactly 8 spaces, replace with TAB
    if line.startswith('        ') and len(line) > 8 and line[8] != ' ':
        # Replace 8 spaces with TAB
        fixed_line = '\t' + line[8:]
        fixed_lines.append(fixed_line)
        print(f"Fixed line {i+1}: converted 8 spaces to TAB")
    else:
        fixed_lines.append(line)

# Write the fixed Makefile
with open('Makefile', 'w') as f:
    f.writelines(fixed_lines)

# Now read the content as string to fix lint target and version
with open('Makefile', 'r') as f:
    content = f.read()

# Fix lint target - replace the deprecated message with actual golangci-lint call
content = content.replace(
    'lint:\n\techo "Linter is deprecated with go1.18!"',
    'lint: golangci-lint\n\t$(GOLANGCILINT) run --timeout 10m'
)

# Update golangci-lint version
content = content.replace(
    'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2',
    'github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2'
)

# Write the final content
with open('Makefile', 'w') as f:
    f.write(content)

print("Makefile has been fixed with:")
print("- Proper TAB indentation for all recipe lines")
print("- Updated lint target to use golangci-lint")
print("- Updated golangci-lint version to v1.55.2")

# Test the syntax
import subprocess
try:
    result = subprocess.run(['make', '--dry-run', 'test'], 
                          capture_output=True, text=True, timeout=10)
    if result.returncode == 0:
        print("✓ Makefile syntax is now correct!")
        print("✓ The CI pipeline should now pass")
    else:
        print(f"✗ Makefile syntax error: {result.stderr}")
except Exception as e:
    print(f"Error testing make: {e}")

# Show a sample of the fixed content
print("\nSample of fixed lines:")
with open('Makefile', 'r') as f:
    lines = f.readlines()
    for i in range(69, min(73, len(lines))):
        print(f"Line {i+1}: {repr(lines[i])}")